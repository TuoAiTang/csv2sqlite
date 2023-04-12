package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Helper struct {
	Name    string
	reader  *csv.Reader
	headers []string
	records [][]string
}

var (
	errEmptyRecords = errors.New("empty records")
)

func NewCSVHelper(name, file string) (*Helper, error) {
	fileInput, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	c := &Helper{
		Name:   name,
		reader: csv.NewReader(fileInput),
	}

	err = c.init()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Helper) init() error {
	all, err := c.reader.ReadAll()
	if err != nil {
		return err
	}

	if len(all) < 2 {
		return errEmptyRecords
	}

	c.headers = all[0]
	c.records = all[1:]
	return nil
}

// TODO : check csv format, especially the length of headers and records
func (c *Helper) check() error {
	return nil
}

func (c *Helper) SaveToDB(dbPath string) error {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	createTableSQL := BuildCreateTableSQL(c.Name, c.headers)
	err = db.Exec(createTableSQL).Error
	if err != nil {
		return err
	}

	err = db.Exec(fmt.Sprintf("delete from %s", c.Name)).Error
	if err != nil {
		return err
	}

	batchInsertSQL, args := BuildBatchInsertSQLWithArgs(c.Name, c.headers, c.records)
	err = db.Exec(batchInsertSQL, args...).Error
	if err != nil {
		return err
	}
	return nil
}

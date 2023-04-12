package csv

import (
	"fmt"
	"strings"
)

func BuildCreateTableSQL(name string, fields []string) string {
	createTableSQL := fmt.Sprintf("create table if not exists %s \n(\n\t", name)
	var fieldDefList []string
	for _, v := range fields {
		fieldDefList = append(fieldDefList, fmt.Sprintf("%s varchar(256) not null", v))
	}
	createTableSQL += strings.Join(fieldDefList, ",\n\t")
	createTableSQL += "\n);"
	return createTableSQL
}

// BuildBatchInsertSQLWithArgs build batch insert sql with args
func BuildBatchInsertSQLWithArgs(name string, fields []string, records [][]string) (string, []interface{}) {
	batchInsertSQL := fmt.Sprintf("insert into %s(%s) values %s", name, strings.Join(fields, ","), strings.TrimRight(strings.Repeat(fmt.Sprintf("(%s),", strings.TrimRight(strings.Repeat("?,", len(fields)), ",")), len(records)), ","))
	var args []interface{}
	for _, record := range records {
		for _, v := range record {
			args = append(args, v)
		}
	}
	return batchInsertSQL, args
}

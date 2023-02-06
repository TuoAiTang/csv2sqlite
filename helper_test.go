package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVHelper_ToDB(t *testing.T) {
	helper, err := NewCSVHelper("coupons", "/Users/xxx/Downloads/coupons.csv")
	assert.Nil(t, err)

	err = helper.SaveToDB("test.db")
	assert.Nil(t, err)
}

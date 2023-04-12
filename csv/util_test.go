package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCreateTableSQL(t *testing.T) {
	type args struct {
		name   string
		fields []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "build",
			args: args{
				name:   "demo",
				fields: []string{"id", "name", "age"},
			},
			want: `create table if not exists demo 
(
	id varchar(256) not null,
	name varchar(256) not null,
	age varchar(256) not null
);`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BuildCreateTableSQL(tt.args.name, tt.args.fields), "BuildCreateTableSQL(%v, %v)", tt.args.name, tt.args.fields)
		})
	}
}

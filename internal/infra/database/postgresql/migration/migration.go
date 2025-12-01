package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{
	&ID011220251300DDLCreateInitialSchema,
}

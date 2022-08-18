package youki_database

import (
	"github.com/go-pg/pg/v9"
	"github.com/AlexZ33/utils"
)

var (
	PostgresTableName = map[string]string{}
)

type PostgresInstance struct {
	db *pg.DB
}

type groupRow struct {
	Name  string `pg:"name"json:"name"`
	Count int    `pg:"count"json:"count"`
}

func (p *PostgresInstance) CreateSchema(schemas iris.Map, citus iris.Map) error {
	db := p.db
	postgres :=
}

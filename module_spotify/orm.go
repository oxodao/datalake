package module_spotify

import "github.com/jmoiron/sqlx"

type ORM struct {
	db *sqlx.DB
}

func newORM(db *sqlx.DB) *ORM {
	return &ORM{db: db}
}

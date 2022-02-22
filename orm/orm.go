package orm

import "github.com/jmoiron/sqlx"

type ORM struct {
	Auth Auth
}

func New(db *sqlx.DB) *ORM {
	return &ORM{
		Auth: Auth{db: db},
	}
}

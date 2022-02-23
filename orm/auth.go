package orm

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/datalake/models"
)

type Auth struct {
	db *sqlx.DB
}

func (a Auth) CreateUser(u *models.User) error {
	_, err := a.db.NamedExec(`
		INSERT INTO datalake_user (name, password)
		VALUES (:name, :password)
	`, u)

	return err
}

func (a Auth) InsertModuleAuthentication(u *models.User, module string, accountId string, data interface{}) error {
	_, err := a.db.NamedExec(`
		INSERT INTO provider_authentication (user_id, module_name, account_username, data)
		VALUES (:user, :module, :account_username, :data)
	`, map[string]interface{}{
		"user":             u.ID,
		"module":           module,
		"account_username": accountId,
		"data":             data,
	})

	return err
}

func (a Auth) FindByUsername(username string) (*models.User, error) {
	rows := a.db.QueryRowx(`
		SELECT id, name, password
		FROM datalake_user
		WHERE name = $1
	`, username)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var user models.User
	err := rows.StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

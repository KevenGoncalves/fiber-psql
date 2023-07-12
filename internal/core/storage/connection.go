package storage

import (
	"database/sql"

	"github.com/KevenGoncalves/fiber-psql/config"
)

func ConnectDB(env config.EnvVars) (*sql.DB, error) {
	db, err := sql.Open(env.DB_DRIVER, env.DB_URI)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}

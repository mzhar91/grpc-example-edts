package config

import (
	"database/sql"
)

type Connection struct {
	Database *sql.DB
}

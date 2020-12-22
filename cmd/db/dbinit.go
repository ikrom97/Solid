package db

import "database/sql"

func DatabaseInit(db *sql.DB) (err error) {
	DDLs := []string{CreateUsersTable}
	for _, ddl := range DDLs {
		_, err = db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return
}

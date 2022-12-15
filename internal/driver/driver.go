// here we connecto to database
package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Capital letter at start -> export function
func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

package utils

import (
	"kurapika/config"
)

func ExecQuery(q string, p []interface{}) error {
	db := config.MysqlDB()

	stmt, err := db.Prepare(q)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(p...)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return nil

}

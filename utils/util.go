package utils

import (
	"kurapika/config"
)

type date struct {
	Day        int
	Week       int
	Month      int
	Year       string
	ActionDate string
}

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

func DateDim(dc date) (dateID int, err error) {
	db := config.MysqlDB()

	stmt, err := db.Prepare("CALL dateDimension(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(dc.Day, dc.Week, dc.Month, dc.Year, dc.ActionDate)
	if err != nil {
		panic(err)
	}

	row, err := db.Query("SELECT @dateId")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&dateID)
	}

	defer db.Close()

	return dateID, nil
}

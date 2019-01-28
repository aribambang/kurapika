package utils

import (
	"fmt"
	"kurapika/config"
	"log"
	"strconv"
	"time"
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

func FetchID(q, ID string) (resID int, err error) {
	db := config.MysqlDB()

	if err := db.QueryRow(q, ID).Scan(&resID); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return resID, nil
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

func Extract(createdAt time.Time) (dc date, err error) {
	_, actionWeek := createdAt.ISOWeek()
	actionYear := createdAt.Year()
	actionMonth := int(createdAt.Month())
	_, _, actionDate := createdAt.Date()

	return date{
		createdAt.YearDay(),
		actionWeek,
		actionMonth,
		strconv.Itoa(actionYear),
		fmt.Sprintf("%d-%d-%d", actionYear, actionMonth, actionDate),
	}, nil
}

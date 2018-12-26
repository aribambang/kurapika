package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func MysqlDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to get env value")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true",
		os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	return db
}

func MgoDB() *mgo.Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to get env value")
	}
	connectionString := fmt.Sprintf("mongodb://%s:27017",
		os.Getenv("MONGO_DB_HOST"),
	)
	dialInfo, err := mgo.ParseURL(connectionString)
	dialInfo.Direct = true
	dialInfo.FailFast = true

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}

	return session
}

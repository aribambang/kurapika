package extractor

import (
	"fmt"
	"kurapika/config"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	CreatedAt time.Time     `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty" json:"updatedAt"`
	Source    string        `bson:"source" json:"source"`
}

func MgoUser() ([]User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to get env value")
	}
	var results []User
	session := config.MgoDB()
	c := session.DB(os.Getenv("API1_DB")).C("user")

	if err := c.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	fmt.Println(results)
	defer session.Close()
	return results, nil
}

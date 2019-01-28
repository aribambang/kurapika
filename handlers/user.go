package handlers

import (
	"fmt"

	"kurapika/extractor"
	"kurapika/transformer"
)

func User() {
	users, err := extractor.MgoUser()
	if err != nil {
		panic(err)
	}
	for i, u := range users {
		fmt.Sprintf("Processing User #%d", i)
		if err := transformer.SQLUser(u); err != nil {
			panic(err)
		}
	}
	fmt.Printf("Done! %d data processed!\n")
}

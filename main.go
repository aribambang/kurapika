package main

import (
	"fmt"
	"time"

	"kurapika/handlers"
)

func main() {

	fmt.Println("Hello World!")

}

func extract() {
	fmt.Println("Kurapika starting")
	start := time.Now().UTC()
	handlers.User()
	fmt.("Done processing in", time.Since(start))
}

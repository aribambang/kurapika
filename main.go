package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"kurapika/handlers"

	"github.com/robfig/cron"
)

func main() {

	fmt.Println("Kurapika loading..")
	c := cron.New()

	c.AddFunc("5 * * * * *", extract)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func extract() {
	fmt.Println("Kurapika starting")
	start := time.Now().UTC()
	handlers.User()
	fmt.Println("Done processing in", time.Since(start))
}

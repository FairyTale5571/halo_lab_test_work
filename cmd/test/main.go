package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/fairytale5571/privat_test/pkg/app"
)

func main() {
	log.Printf("start application\n")
	a, err := app.New()
	if err != nil {
		log.Fatal("start application failed: ", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	if err := a.DB.Close(); err != nil {
		log.Fatal("close database failed: ", err)
		return
	}
	a.Logger.Info("shutdown application")
}

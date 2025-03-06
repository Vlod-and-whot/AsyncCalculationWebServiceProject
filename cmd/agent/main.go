package main

import (
	"log"

	"github.com/Vlod-and-whot/AsyncCalculationWebServiceProject/internal/application"
)

func main() {
	app := application.NewAgent()
	log.Println("Starting Agent")
	app.Run()
}


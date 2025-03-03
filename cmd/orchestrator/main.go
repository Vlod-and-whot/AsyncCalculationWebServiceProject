package main

import (
	"log"

	"github.com/Vlod-and-whot/AsyncCalculationWebServiceProject/internal/application"
)

func main() {
	app := application.NewOrchestrator()
	log.Println("Starting Orchestrator on port", app.Config.Addr)
	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}
}

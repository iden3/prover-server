package main

import (
	"os"

	"github.com/iden3/prover-server/pkg/app"
	"github.com/iden3/prover-server/pkg/app/configs"
	"github.com/iden3/prover-server/pkg/app/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	config, err := configs.ReadConfigFromFile("prover")
	if err != nil {
		log.Error("cannot read issuer config storage", err.Error())
		os.Exit(1)
	}

	// init handlers for router

	var appHandlers = app.Handlers{
		ZKHandler: handlers.NewZKHandler(config.Prover),
	}
	router := appHandlers.Routes()

	server := app.NewServer(router)

	// start the server
	server.Run(config.Server.Port)

}

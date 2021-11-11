package main

import (
	"context"
	"github.com/iden3/prover-server/pkg/app"
	"github.com/iden3/prover-server/pkg/app/configs"
	"github.com/iden3/prover-server/pkg/app/handlers"
	"github.com/iden3/prover-server/pkg/log"
	"go.uber.org/zap"
	"os"
)

func main() {

	config, err := configs.ReadConfigFromFile("prover")
	if err != nil {
		log.Error(context.Background(), "cannot read issuer config storage", zap.Error(err))
		os.Exit(1)
	}

	log.SetLevelStr("debug")
	// init handlers for router

	var appHandlers = app.Handlers{
		ZKHandler: handlers.NewZKHandler(config.Prover),
	}
	router := appHandlers.Routes()

	server := app.NewServer(router)

	// start the server
	server.Run(config.Server.Port)

}

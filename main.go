package main

import (
	"assignment-permission/config"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	var err error
	if config.LoadConfig(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ctx := context.Background()
	err = initDB(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Info("Database connected 🚀")
	}

	StartServer()
}

package main

import (
	"GAMELIB/internal/server"
	"GAMELIB/internal/storage/minio"
	"GAMELIB/internal/storage/postgres"
	"GAMELIB/pkg/config"
	"GAMELIB/pkg/web"
	"flag"
	"log"
)

func main() {
	env := flag.String("env", "dev", "env for start")
	flag.Parse()

	if err := config.ParseConfig(*env); err != nil {
		log.Fatalf("[Main] Error initialize configs: %s", err.Error())
	}

	cfg := &config.Config{
		Server:  web.InitServerConfig(),
		Storage: postgres.InitStorageConfig(),
		Minio:   minio.InitMinioConfig(),
	}

	app, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("[Main] Can't connect to database: %s", err.Error())
	} else {
		log.Println("Database connection!")
	}

	path := cfg.Server.Port
	if err := app.Start(path); err != nil {
		log.Fatalf("[Main] Server start error: %s", err.Error())
	}
}

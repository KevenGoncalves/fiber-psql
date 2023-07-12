package main

import (
	"github.com/KevenGoncalves/fiber-psql/config"
	"github.com/KevenGoncalves/fiber-psql/internal/core/server"
	"github.com/KevenGoncalves/fiber-psql/pkg/shutdown"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	//load config
	env, err := config.LoadConfig()
	if err != nil {
		log.Printf("error: %v", err)
		exitCode = 1
		return
	}

	//run server
	cleanup, err := run(env)

	// run cleanup after server terminate
	defer cleanup()

	if err != nil {
		log.Printf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := server.BuildServer(env)

	if err != nil {
		return nil, err
	}

	//start server
	go func() {
		app.Listen(env.SERVER_PORT)
	}()

	// return cleanup to close database and server connection
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

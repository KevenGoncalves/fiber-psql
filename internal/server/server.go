package server

import (
	"github.com/KevenGoncalves/fiber-psql/config"
	"github.com/KevenGoncalves/fiber-psql/internal/app/users"
	"github.com/KevenGoncalves/fiber-psql/internal/databases/postgres"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BuildServer(env config.EnvVars) (*fiber.App, func(), error) {

	//connect database
	db, err := postgres.ConnectDB(env)
	if err != nil {
		return nil, nil, err
	}

	//create app
	app := fiber.New(fiber.Config{
		JSONDecoder: sonic.Unmarshal,
		JSONEncoder: sonic.Marshal,
	})

	//middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	//create Domain
	users.Routes(app, users.NewUserController(db))

	return app, func() {
		postgres.CloseDB(db)
	}, nil
}

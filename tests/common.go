package tests

import (
	"GoDeveloperVK-testTask/auth"
	"GoDeveloperVK-testTask/server/handler"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/database"
	"GoDeveloperVK-testTask/utils/logger"
	"errors"
	"log/slog"
)

var Config utils.Config = utils.Config{
	Env:        "dev",
	HTTPServer: utils.HTTPServer{Host: "localhost", Port: "5432"},
	Database:   utils.Database{Host: "localhost", Port: "5432", Username: "postgres", Password: "postgres", DBName: "postgres", SSLMode: "disable"},
	Auth:       utils.Auth{SecretKey: "secret_key", Salt: "salt"},
}

func StartDB() {
	_, err := database.InitDatabase(
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.DBName,
		Config.Database.SSLMode,
	)

	if err != nil {
		logger.Err(errors.New("failed to init database"))
	}
}

func StartApp() {
	log := logger.SetupLogger(Config.Env)
	log.Info("You are in mode:", slog.String("env", Config.Env))

	_, err := database.InitDatabase(
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.DBName,
		Config.Database.SSLMode,
	)
	if err != nil {
		log.Error("failed to init database", logger.Err(err))
	}

	auth.InitAuth(&Config.Auth)

	handler.InitActorsHandlers()
	handler.InitMoviesHandlers()
}

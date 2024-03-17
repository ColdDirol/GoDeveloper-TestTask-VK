package main

import (
	"GoDeveloperVK-testTask/auth"
	"GoDeveloperVK-testTask/server/handler"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/database"
	"GoDeveloperVK-testTask/utils/logger"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	config := utils.InitConfig("config.json")

	log := logger.SetupLogger(config.Env)
	log.Info("You are in mode:", slog.String("env", config.Env))

	db, err := database.InitDatabase(
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)
	if err != nil {
		log.Error("failed to init database", logger.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	auth.InitAuth(&config.Auth)
	handler.InitActorsHandlers()
	handler.InitMoviesHandlers()

	runApplication(&config.HTTPServer, log)
}

func runApplication(httpConfig *utils.HTTPServer, log *slog.Logger) {
	log.Info("Server start listening on port: ", slog.String("port", httpConfig.Port))
	httpAddress := httpConfig.Host + ":" + httpConfig.Port
	err := http.ListenAndServe(httpAddress, nil)
	if err != nil {
		log.Error("Failed to start server", logger.Err(err))
		os.Exit(1)
	}
}

package main

import (
	"net/http"

	"financial-app-backend/internal/database"

	_ "github.com/lib/pq" // here
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"
	"financial-app-backend/internal/api"
	"financial-app-backend/internal/config"
)

//Create Server object and start listener
func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting server.")
	//Creating new database
	db, err := database.New()
	if err != nil {
		logrus.WithError(err).Fatal("Error verifying database.")
	}
	//Creating new router
	router, err := api.NewRouter(db)
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	//Starting server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed.")
	}
}

//susu

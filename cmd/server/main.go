package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ednardomartins/gerenciador-financeiro-transporte/domain/transporte"
	"github.com/ednardomartins/gerenciador-financeiro-transporte/internal/server/http"
)

const (
	envWorkshopServicePort = "WORKSHOP_SERVICE_PORT"
	envDatabaseName        = "DATABASE_NAME"
	envMongoURL            = "MONGO_URL"

	defaultDatabaseName = "workshop"
	defaultPort         = "8080"
)

func main() {
	/*
	 * Services...
	 */
	transporteService := transporte.NewService()
	/*
	 * Handler...
	 */
	handler := http.NewHandler(transporteService)
	/*
	 * Server...
	 */
	server := http.New(getApplicationPort(), handler)
	server.ListenAndServe()
	/*
	 * Graceful shutdown...
	 */
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}

func getApplicationPort() string {
	return getEnvVar(envWorkshopServicePort, defaultPort)
}

func getEnvVar(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return value
}

package main

import (
	"fmt"
	consulServiceManager "github.com/ayushs-2k4/go-consul-service-manager"
	"kontest-user-stats-service/routes"
	"kontest-user-stats-service/utils"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	applicationHost = "localhost"                  // Default value for local development
	applicationPort = "5151"                       // Default value for local development
	serviceName     = "KONTEST-USER-STATS-SERVICE" // Service name for Service Registry
	consulHost      = "localhost"                  // Default value for local development
	consulPort      = 5150                         // Port as a constant (can be constant if it won't change)
)

func initializeVariables() {
	// Attempt to read the KONTEST_API_USER_STATS_SERVICE_HOST environment variable
	if host := os.Getenv("KONTEST_API_USER_STATS_SERVICE_HOST"); host != "" {
		applicationHost = host // Override with the environment variable if set
	}

	// Attempt to read the KONTEST_API_SERVER_PORT environment variable
	if port := os.Getenv("KONTEST_API_USER_STATS_SERVICE_PORT"); port != "" {
		applicationPort = port // Override with the environment variable if set
	}

	// Attempt to read the CONSUL_ADDRESS environment variable
	if host := os.Getenv("CONSUL_HOST"); host != "" {
		consulHost = host // Override with the environment variable if set
	}

	// Attempt to read the CONSUL_PORT environment variable
	if port := os.Getenv("CONSUL_PORT"); port != "" {
		if portInt, err := strconv.Atoi(port); err == nil {
			consulPort = portInt // Override with the environment variable if set and valid
		}
	}
}

func main() {
	initializeVariables()

	utils.InitializeDependencies()

	portInt, err := strconv.Atoi(applicationPort)
	if err != nil {
		log.Fatalf("Failed to convert port to integer: %v", err)
	}

	consulService := consulServiceManager.NewConsulService(consulHost, consulPort)
	consulService.Start(applicationHost, portInt, serviceName)

	router := http.NewServeMux()

	routes.RegisterRoutes(router)

	server := http.Server{
		Addr:    ":" + applicationPort, // Use the field name Addr for the address
		Handler: router,                // Use the field name Handler for the router
	}

	fmt.Println("Server listening at port: " + applicationPort)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	consulServiceManager "github.com/ayushs-2k4/go-consul-service-manager/consulservicemanager"
	"kontest-user-stats-service/routes"
	"kontest-user-stats-service/utils"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	applicationHost = "localhost"                  // Default value for local development
	applicationPort = 5151                         // Default value for local development
	serviceName     = "KONTEST-USER-STATS-SERVICE" // Service name for Service Registry
	consulHost      = "localhost"                  // Default value for local development
	consulPort      = 5150                         // Port as a constant (can be constant if it won't change)
)

func initializeVariables() {
	// Get the hostname of the machine
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error fetching hostname: %v", err)
	}

	// Attempt to read the KONTEST_API_USER_STATS_SERVICE_HOST environment variable
	if host := os.Getenv("KONTEST_API_USER_STATS_SERVICE_HOST"); host != "" {
		applicationHost = host // Override with the environment variable if set
	} else {
		applicationHost = hostname // Use the machine's hostname if the env var is not set
	}

	// Attempt to read the KONTEST_API_SERVER_PORT environment variable
	if port := os.Getenv("KONTEST_API_USER_STATS_SERVICE_PORT"); port != "" {
		parsedPort, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("Invalid port value: %v", err)
		}
		applicationPort = parsedPort // Override with the environment variable if set
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

	consulService := consulServiceManager.NewConsulService(consulHost, consulPort)
	consulService.Start(applicationHost, applicationPort, serviceName, []string{})

	router := http.NewServeMux()

	routes.RegisterRoutes(router)

	server := http.Server{
		Addr:    ":" + strconv.Itoa(applicationPort), // Use the field name Addr for the address
		Handler: router,                              // Use the field name Handler for the router
	}

	fmt.Println("Server listening at port: " + strconv.Itoa(applicationPort))

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

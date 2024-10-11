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

var serviceName = "KONTEST-USER-STATS-SERVICE"

func main() {
	utils.InitializeDependencies()

	port := os.Getenv("KONTEST_API_USER_STATS_SERVICE_PORT")
	if port == "" {
		port = "5152"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to convert port to integer: %v", err)
	}

	consulService := consulServiceManager.NewConsulService("localhost", 5150)
	consulService.Start(portInt, serviceName)

	router := http.NewServeMux()

	routes.RegisterRoutes(router)

	server := http.Server{
		Addr:    ":" + port, // Use the field name Addr for the address
		Handler: router,     // Use the field name Handler for the router
	}

	fmt.Println("Server listening at port: " + port)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

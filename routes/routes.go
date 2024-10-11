package routes

import (
	"fmt"
	"kontest-user-stats-service/controllers"
	"net/http"
)

func HelloGETHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! GET")
}

func HelloPOSTHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! POST")
}

func HelloPUTHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! PUT")
}

func HelloDELETEHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! DELETE")
}

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /codechef_user", controllers.GetCodechefUser)

	registerHelloRoutes(router)
}

func registerHelloRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /hello", HelloGETHandler)
	router.HandleFunc("POST /hello", HelloPOSTHandler)
	router.HandleFunc("DELETE /hello", HelloDELETEHandler)
	router.HandleFunc("PUT /hello", HelloPUTHandler)
}

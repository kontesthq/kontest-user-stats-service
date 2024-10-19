package controllers

import (
	"encoding/json"
	"fmt"
	"kontest-user-stats-service/utils"
	"log/slog"
	"net/http"
)

func GetCodechefUser(w http.ResponseWriter, r *http.Request) {
	// Get the username from the URL
	username := r.URL.Query().Get("username")
	if username == "" {
		slog.Error(fmt.Sprintf("Username not provided"))
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	codechefService := utils.GetDependencies().CodechefService
	codechefUser, err := codechefService.GetUserData(username)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get user data: %v", err))
		http.Error(w, "Failed to get user data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user data in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codechefUser)
}

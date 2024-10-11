package controllers

import (
	"encoding/json"
	"kontest-user-stats-service/utils"
	"net/http"
)

func GetCodechefUser(w http.ResponseWriter, r *http.Request) {
	// Get the username from the URL
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	codechefService := utils.GetDependencies().CodechefService
	codechefUser, err := codechefService.GetUserData(username)
	if err != nil {
		http.Error(w, "Failed to get user data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user data in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codechefUser)
}

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"kontest-user-stats-service/exceptions"
	"kontest-user-stats-service/model/codechef"
	"log"
	"net/http"
)

type CodeChefService struct {
	hTTPClient *http.Client
}

func NewCodeChefService() *CodeChefService {
	return &CodeChefService{
		hTTPClient: &http.Client{},
	}
}

func (s *CodeChefService) GetUserData(username string) (*codechef.CodeChefUser, error) {
	mainURL := "https://codechef-api.vercel.app/handle/"
	url := mainURL + username

	resp, err := s.hTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user data: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var codeChefUser codechef.CodeChefUser
	if err := json.Unmarshal(bodyBytes, &codeChefUser); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if codeChefUser.Success {
		return &codeChefUser, nil
	}
	return nil, &exceptions.CodeChefException{
		Message:   "Username not valid",
		ErrorType: exceptions.UsernameNotFound,
	}
}

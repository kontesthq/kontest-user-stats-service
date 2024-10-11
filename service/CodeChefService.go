package service

import (
	"encoding/json"
	"fmt"
	"io"
	"kontest-user-stats-service/exceptions"
	"kontest-user-stats-service/model/codechef"
	custom_marshals "kontest-user-stats-service/model/custom-marshals"
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

func (s *CodeChefService) GetUserKontests(username string) ([]codechef.RatingDataEntry, error) {
	url := "https://codechef-api.vercel.app/handle/" + username

	resp, err := s.hTTPClient.Get(url)
	if err != nil {
		log.Printf("Error in downloading CodeChef profile: %v", err)
		return nil, &exceptions.CodeChefException{
			Message:   "Failed to fetch user data",
			ErrorType: exceptions.UsernameNotFound,
		}
	}
	defer resp.Body.Close()

	var rawHTML string
	if err := json.NewDecoder(resp.Body).Decode(&rawHTML); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return s.GetUserDataFromRawHTML(rawHTML)
}

func (s *CodeChefService) GetUserDataFromRawHTML(stringHTML string) ([]codechef.RatingDataEntry, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(stringHTML), &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal HTML string: %w", err)
	}

	ratings, ok := obj["ratingData"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("no ratingData found")
	}

	var contests []codechef.RatingDataEntry
	for _, entry := range ratings {
		ratingEntry, ok := entry.(map[string]interface{})
		if !ok {
			continue // Skip if the entry is not a map
		}

		// Create a new contest entry with default values
		contest := codechef.RatingDataEntry{
			Code:        ratingEntry["code"].(string),
			Reason:      nil, // Initialize to nil
			PenalisedIn: nil, // Use pointer to handle null values
			Name:        ratingEntry["name"].(string),
			EndDate:     ratingEntry["end_date"].(string),
			Color:       ratingEntry["color"].(string),
		}

		// Check for optional fields and assign if they exist
		if year, exists := ratingEntry["getyear"]; exists {
			if y, ok := year.(float64); ok {
				yearInt := int(y)
				contest.Year = custom_marshals.IntString(yearInt) // Assign year
			}
		}

		if month, exists := ratingEntry["getmonth"]; exists {
			if m, ok := month.(float64); ok {
				monthInt := int(m)
				contest.Month = custom_marshals.IntString((monthInt)) // Assign month
			}
		}

		if day, exists := ratingEntry["getday"]; exists {
			if d, ok := day.(float64); ok {
				dayInt := int(d)
				contest.Day = custom_marshals.IntString((dayInt)) // Assign day
			}
		}

		if reason, exists := ratingEntry["reason"]; exists {
			if r, ok := reason.(string); ok {
				contest.Reason = &r // Assign reason
			}
		}

		if penalisedIn, exists := ratingEntry["penalised_in"]; exists {
			if p, ok := penalisedIn.(bool); ok {
				contest.PenalisedIn = &p // Assign penalisedIn as pointer
			}
		}

		if rating, exists := ratingEntry["rating"]; exists {
			if r, ok := rating.(float64); ok {
				ratingInt := int(r)
				contest.Rating = custom_marshals.IntString(ratingInt) // Assign rating
			}
		}

		if rank, exists := ratingEntry["rank"]; exists {
			if r, ok := rank.(float64); ok {
				rankInt := int(r)
				contest.Rank = custom_marshals.IntString(rankInt) // Assign rank
			}
		}

		// Append the constructed contest to the contests slice
		contests = append(contests, contest)
	}

	return contests, nil
}

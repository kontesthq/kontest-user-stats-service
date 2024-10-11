package codechef

import custom_marshals "kontest-user-stats-service/model/custom-marshals"

type RatingDataEntry struct {
	Code        string                    `json:"code"`
	Year        custom_marshals.IntString `json:"getyear"`
	Month       custom_marshals.IntString `json:"getmonth"`
	Day         custom_marshals.IntString `json:"getday"`
	Reason      *string                   `json:"reason"`       // Use pointer to handle null values
	PenalisedIn *bool                     `json:"penalised_in"` // Change to *bool to handle null values
	Rating      custom_marshals.IntString `json:"rating"`
	Rank        custom_marshals.IntString `json:"rank"`
	Name        string                    `json:"name"`
	EndDate     string                    `json:"end_date"`
	Color       string                    `json:"color"`
}

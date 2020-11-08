package politicker

import (
	"encoding/json"
	"net/http"
	"os"
)

// Response is the response from propublica api
type Response struct {
	Results []Result `json:"results"`
}

// Result is the unfortunate side effect of the weird format of the API data
type Result struct {
	Congress   string   `json:"congress"`
	Chamber    string   `json:"chamber"`
	NumResults int      `json:"num_results"`
	Offset     int      `json:"offset"`
	Members    []Member `json:"members"`
}

// Member represents a member of Congress
type Member struct {
	ID                   string      `json:"id"`
	Title                string      `json:"title"`
	ShortTitle           string      `json:"short_title"`
	APIURI               string      `json:"api_uri"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	DateOfBirth          string      `json:"date_of_birth"`
	Gender               string      `json:"gender"`
	Party                string      `json:"party"`
	LeadershipRole       string      `json:"leadership_role"`
	URL                  string      `json:"url"`
	RssURL               string      `json:"rss_url"`
	ContactForm          string      `json:"contact_form"`
	InOffice             bool        `json:"in_office"`
	DwNominate           float64     `json:"dw_nominate"`
	IdealPoint           interface{} `json:"ideal_point"`
	Seniority            string      `json:"seniority"`
	NextElection         string      `json:"next_election"`
	TotalVotes           int         `json:"total_votes"`
	MissedVotes          int         `json:"missed_votes"`
	TotalPresent         int         `json:"total_present"`
	LastUpdated          string      `json:"last_updated"`
	Office               string      `json:"office"`
	Phone                string      `json:"phone"`
	Fax                  string      `json:"fax"`
	State                string      `json:"state"`
	SenateClass          string      `json:"senate_class"`
	StateRank            string      `json:"state_rank"`
	MissedVotesPct       float64     `json:"missed_votes_pct"`
	VotesWithPartyPct    float64     `json:"votes_with_party_pct"`
	VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
}

// TODO: Make this take an interface so we can decode other things besides members
func (app *appEnv) fetchJSON(url string, data *Response) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", os.Getenv("PROPUBLICA_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

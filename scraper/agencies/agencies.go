package agencies

import (
	"github.com/b-eau/ecfr-analyzer/scraper/fetch_json"
)

type AgenciesResponse struct {
	Agencies []Agency `json:"agencies"`
}

type Agency struct {
	Name          string         `json:"name"`
	ShortName     string         `json:"short_name"`
	DisplayName   string         `json:"display_name"`
	SortableName  string         `json:"sortable_name"`
	Slug          string         `json:"slug"`
	Children      []Agency       `json:"children"`
	CFRReferences []CFRReference `json:"cfr_references"`
}

type CFRReference struct {
	Title   int    `json:"title"`
	Chapter string `json:"chapter"`
}

func agenciesUrl() string {
	return "https://www.ecfr.gov/api/admin/v1/agencies.json"
}

func FetchAgencies() (AgenciesResponse, error) {
	var response AgenciesResponse
	err := fetch_json.FetchJson(agenciesUrl(), &response)
	return response, err
}

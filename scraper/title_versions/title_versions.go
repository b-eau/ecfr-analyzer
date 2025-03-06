package title_versions

import (
	"fmt"

	"github.com/b-eau/ecfr-analyzer/scraper/fetch_json"
)

type TitleContentVersions struct {
	ContentVersions []TitleContentVersion `json:"content_versions"`
	Meta            Meta                  `json:"meta"`
}

type TitleContentVersion struct {
	Date          string  `json:"date"`
	AmendmentDate string  `json:"amendment_date"`
	IssueDate     string  `json:"issue_date"`
	Identifier    string  `json:"identifier"`
	Name          string  `json:"name"`
	Part          string  `json:"part"`
	Substantive   bool    `json:"substantive"`
	Removed       bool    `json:"removed"`
	Subpart       *string `json:"subpart"` // Can be string or null
	Title         string  `json:"title"`
	Type          string  `json:"type"`
}

type Meta struct {
	Title               string     `json:"title"`
	ResultCount         string     `json:"result_count"`
	LatestAmendmentDate string     `json:"latest_amendment_date"`
	LatestIssueDate     string     `json:"latest_issue_date"`
	IssueDate           *DateRange `json:"issue_date,omitempty"` // Optional field
}

type DateRange struct {
	Lte string `json:"lte"`
	Gte string `json:"gte"`
}

func titleVersionsUrl(titleId string) string {
	return fmt.Sprintf("https://www.ecfr.gov/api/versioner/v1/versions/title-%v.json", titleId)
}

func FetchTitleVersions(titleId string) (TitleContentVersions, error) {
	var titleVersions TitleContentVersions
	err := fetch_json.FetchJson(titleVersionsUrl(titleId), &titleVersions)
	return titleVersions, err
}

package title_summaries

import (
	"github.com/b-eau/ecfr-analyzer/scraper/fetch_json"
)

type TitleSummariesResponse struct {
	Titles []TitleSummary   `json:"titles"`
	Meta   TitleSummaryMeta `json:"meta"`
}

type TitleSummary struct {
	Number               int    `json:"number"`
	Name                 string `json:"name"`
	LatestAmendedOn      string `json:"latest_amended_on"`
	LatestIssueDate      string `json:"latest_issue_date"`
	UpToDateAsOf         string `json:"up_to_date_as_of"`
	Reserved             bool   `json:"reserved"`
	ProcessingInProgress bool   `json:"processing_in_progress,omitempty"` //added omitempty because it's only present sometimes.
}

type TitleSummaryMeta struct {
	Date             string `json:"date"`
	ImportInProgress bool   `json:"import_in_progress"`
}

func listTitleSummariesUrl() string {
	return "https://www.ecfr.gov/api/versioner/v1/titles.json"
}

func FetchTitleSummaries() (TitleSummariesResponse, error) {
	var response TitleSummariesResponse
	err := fetch_json.FetchJson(listTitleSummariesUrl(), &response)
	return response, err
}

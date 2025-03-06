package main

import (
	"fmt"
	"log"

	"github.com/b-eau/ecfr-analyzer/scraper/agencies"
	"github.com/b-eau/ecfr-analyzer/scraper/title_summaries"
	"github.com/b-eau/ecfr-analyzer/scraper/title_versions"
	"github.com/b-eau/ecfr-analyzer/scraper/write_json"
)

func dataFilePath(fileName string) string {
	return fmt.Sprintf("data/%v", fileName)
}

func main() {
	agenciesResponse, err := agencies.FetchAgencies()
	if err != nil {
		log.Fatalf("error fetching agencies: %v", err)
	}
	if err = write_json.WriteJSONToFile(dataFilePath("agencies.json"), agenciesResponse); err != nil {
		log.Fatalf("error writing agencies: %v", err)
	}
	titleSummariesResponse, err := title_summaries.FetchTitleSummaries()
	if err != nil {
		log.Fatalf("error fetching title summaries: %v", err)
	}
	if err := write_json.WriteJSONToFile(dataFilePath("title_summaries.json"), titleSummariesResponse); err != nil {
		log.Fatalf("error writing title summaries: %v", err)
	}
	for _, title := range titleSummariesResponse.Titles {
		titleVersionsResponse, err := title_versions.FetchTitleVersions(fmt.Sprintf("%v", title.Number))
		if err != nil {
			log.Fatalf("error fetching title versions: %v", err)
		}
		filePath := dataFilePath(fmt.Sprintf("title_%v_versions.json", title.Number))
		if err := write_json.WriteJSONToFile(filePath, titleVersionsResponse); err != nil {
			log.Fatalf("error writing title versions: %v", err)
		}
	}
}

package domain

import (
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
)

func TestStatistics(t *testing.T) {
	docTest, err := download.Downloader(config.StatisticsPlayerGoalScoredURL)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(docTest.Html())
	tests := []struct {
		doc *goquery.Document
	}{
		{
			doc: docTest,
		},
	}
	for _, test := range tests {
		PlayerStatisticsWithGoalsScored(test.doc)
	}
}

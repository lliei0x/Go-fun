package domain

import (
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
)

func TestTeams(t *testing.T) {
	docTest, err := download.Downloader(config.TeamsURL)
	if err != nil {
		log.Println(err)
	}
	tests := []struct {
		doc *goquery.Document
	}{
		{
			doc: docTest,
		},
	}

	for _, test := range tests {
		Teams(test.doc)
	}
}

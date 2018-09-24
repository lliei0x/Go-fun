package domain

import (
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
)

func TestMateches(t *testing.T) {
	docTest, err := download.Downloader(config.MatchesURL)
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
		Matches(test.doc)
	}
}

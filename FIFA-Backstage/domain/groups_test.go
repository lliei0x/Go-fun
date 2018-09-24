package domain

import (
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
)

func TestGroups(t *testing.T) {
	docTest, err := download.Downloader(config.GroupsURL)
	if err != nil {
		log.Panicln(err)
	}
	tests := []struct {
		doc *goquery.Document
	}{
		{
			doc: docTest,
		},
	}
	for _, test := range tests {
		Groups(test.doc)
	}
}

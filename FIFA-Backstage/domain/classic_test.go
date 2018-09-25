package domain

import (
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
)

func TestClassic(t *testing.T) {
	docTest, err := download.Downloader(config.ClassicURL)
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
		Classic(test.doc)
	}
}

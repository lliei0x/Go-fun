package domain

import (
	"log"
	"testing"

	"leeif.me/Go-fun/FIFA-Backstage/infra/download"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
)

func TestPlayers(t *testing.T) {
	docTest, err := download.Downloader(config.PlayersURLList)
	// src, _ := docTest.Html()
	// fmt.Println(src)
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
		Players(test.doc)
	}
}

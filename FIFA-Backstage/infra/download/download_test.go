package download

import (
	"testing"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: "https://www.fifa.com/worldcup/matches/",
		},
	}
	for _, test := range tests {
		doc, err := Downloader(test.url)
		newDoc, _ := doc.Html()
		t.Log(newDoc, err)
	}
}

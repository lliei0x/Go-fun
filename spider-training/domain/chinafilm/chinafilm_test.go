package chinafilm

import (
	"testing"

	"leeif.me/Go-fun/spider-training/infra/url"
)

func TestGetChinaFilmResponse(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.ChinaFilmRealURL,
		},
	}

	for _, test := range tests {
		GetChinaFilmResponse(test.url)
	}
}

func TestGetChinaFilmMethodTwo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.ChinaFilmRealURL,
		},
	}

	for _, test := range tests {
		GetChinaFilmMethodTwo(test.url)
	}
}

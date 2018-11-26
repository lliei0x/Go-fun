package downloader

import (
	"testing"
)

func TestGetHttpResponse(t *testing.T) {
	tests := []struct {
		url string
		ok  bool
	}{
		{
			url: "http://www.piaofang.biz/",
			ok:  true,
		},
		{
			url: "http://www.piaofang.biz/",
			ok:  false,
		},
	}
	for _, test := range tests {
		responseByte, err := GetHttpResponse(test.url, test.ok)
		if err != nil {
			t.Error(err)
		}
		t.Log(string([]byte(responseByte)))
	}
}

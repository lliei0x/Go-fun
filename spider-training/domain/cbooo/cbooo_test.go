package cbooo

import (
	"testing"

	"leeif.me/Go-fun/spider-training/infra/url"
)

func TestGetRealTimeRankInfo(t *testing.T) {
	tests := []struct {
		url string
		ok  bool
	}{
		{
			url: url.CboRootURL,
			ok:  false,
		},
	}
	for _, test := range tests {
		realTimeRankInfo, err := GetRealTimeRankInfo(test.url, test.ok)
		if err != nil {
			t.Error(err)
		}
		t.Log(realTimeRankInfo)
	}
}

package piaofang

import (
	"testing"
)

func TestGetPiaoFangRankInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: "http://www.piaofang.biz/",
		},
	}
	for _, test := range tests {
		GetPiaofangRankInfo(test.url)
	}
}

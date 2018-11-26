package cbooo

import (
	"strings"

	"leeif.me/Go-fun/spider-training/infra/downloader"
	"leeif.me/Go-fun/spider-training/infra/errors"

	"github.com/PuerkitoBio/goquery"
)

type RealTimeRankInfo struct {
	Rank              string `json:"rank"`
	MovieName         string `json="movie_name"`
	RealTimeBoxOffice string `json="real_time_box_office"`
	BoxOfficeRatio    string `json="box_office_iratio"`
	BoxOfficeIncome   string `json="box_office_income"`
	TicketRate        string `json:"ticket_rate"`
	ReleaseDate       string `json="release_date"`
}

func GetRealTimeRankInfo(url string, ok bool) ([]RealTimeRankInfo, error) {
	responseByte, err := downloader.GetHttpResponse(url, ok)
	if err != nil {
		return nil, errors.ErrorCboooRealTimeRankInfo
	}
	responseString := string([]byte(responseByte))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	if err != nil {
		return nil, errors.ErrorNewDocument
	}
	var realTimeRankInfoList []RealTimeRankInfo
	doc.Find("#topdatatr tr").Each(func(i int, selection *goquery.Selection) {
		realTimeRankInfo := RealTimeRankInfo{
			Rank:              selection.Find("td").Eq(0).Text(),
			MovieName:         selection.Find("td").Eq(1).Text(),
			RealTimeBoxOffice: selection.Find("td").Eq(2).Text(),
			BoxOfficeRatio:    selection.Find("td").Eq(3).Text(),
			BoxOfficeIncome:   selection.Find("td").Eq(4).Text(),
			TicketRate:        selection.Find("td").Eq(5).Text(),
			ReleaseDate:       selection.Find("td").Eq(6).Text(),
		}
		realTimeRankInfoList = append(realTimeRankInfoList, realTimeRankInfo)
	})
	return realTimeRankInfoList, nil
}

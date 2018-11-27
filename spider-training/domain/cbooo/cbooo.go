package cbooo

import (
	"fmt"
	"strconv"
	"strings"

	"leeif.me/Go-fun/spider-training/infra/downloader"
	"leeif.me/Go-fun/spider-training/infra/errors"
	urlList "leeif.me/Go-fun/spider-training/infra/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
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

type MovieListCBOInfo struct {
	BoxOffice    string `json:"box_office"`
	MovieID      string `json:"movie_id"`
	MovieEnName  string `json:"movie_en_name"`
	MovieName    string `json:"movie_name"`
	Ranking      string `json:"ranking"`
	DefaultImage string `json:"default_image"`
	ReleaseYear  string `json:"release_year"`
	Country      string `json:"country"`
}

var (
	areaList      []int
	totalPageList []int64
	testArea      = 5
)

func GetMovieListCBOInfo(url string, ok bool) {
	getTotalPage(url, ok)
	for index := range totalPageList {
		for i := 1; i <= int(totalPageList[index]); i++ {
			realURL := fmt.Sprintf(urlList.CboRealURL, areaList[testArea], i)
			responseByte, err := downloader.GetHttpResponse(realURL, ok)
			if err != nil {
				return
			}
			responseString := string([]byte(responseByte))
			info := gjson.Parse(responseString).Get("pData")

			for _, oneInfo := range info.Array() {
				var oneRank MovieListCBOInfo
				oneRank = MovieListCBOInfo{
					Ranking:      oneInfo.Get("Ranking").String(),
					MovieID:      oneInfo.Get("ID").String(),
					MovieEnName:  oneInfo.Get("MovieEnName").String(),
					MovieName:    oneInfo.Get("MovieName").String(),
					ReleaseYear:  oneInfo.Get("releaseYear").String(),
					DefaultImage: oneInfo.Get("defaultImage").String(),
					BoxOffice:    oneInfo.Get("BoxOffice").String(),
					Country:      string(areaList[testArea]),
				}
				fmt.Println(oneRank)
				// initial.DataBase.Create(&oneRank)
			}
		}
	}
}

func getTotalPage(url string, ok bool) {
	getMovieListOfArea(url, ok)
	var total int64
	for i := range areaList[testArea-1 : testArea] {
		oneURL := fmt.Sprintf(urlList.CboRealURL, areaList[testArea], i+1)
		responseByte, err := downloader.GetHttpResponse(oneURL, false)
		if err != nil {
			return
		}
		responseString := string([]byte(responseByte))
		totalPage := gjson.Parse(responseString).Get("tPage").Int()
		total = gjson.Parse(responseString).Get("tCount").Int()
		totalPageList = append(totalPageList, totalPage)
	}
	for i := range totalPageList {
		fmt.Printf("国家%v，总共%v页，有%v条电影数据", areaList[testArea], totalPageList[i], total)
		fmt.Println("")
	}
}

func getMovieListOfArea(url string, ok bool) {
	responseByte, err := downloader.GetHttpResponse(url, ok)
	if err != nil {
		return
	}
	responseString := string([]byte(responseByte))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(responseString))
	if err != nil {
		return
	}
	doc.Find("#selArea option").Each(func(i int, selection *goquery.Selection) {
		// area := selection.Text()
		areaString, _ := selection.Attr("value")
		areaID, _ := strconv.Atoi(areaString)
		// fmt.Println(areaID, area)
		areaList = append(areaList, areaID)
	})
	fmt.Println(len(areaList))
	// doc.Find("#selType option").Each(func(i int, selection *goquery.Selection) {
	// 	category := selection.Text()
	// 	categoryString, _ := selection.Attr("value")
	// 	categoryID, _ := strconv.Atoi(categoryString)
	// 	fmt.Println(categoryID, category)
	// 	categoryList = append(categoryList, categoryID)
	// })

	// return areaList, nil
}

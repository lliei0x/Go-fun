package piaofang

import (
	"fmt"
	"regexp"

	"leeif.me/Go-fun/spider-training/library/utils/adpater"

	"leeif.me/Go-fun/spider-training/infra/downloader"
	"leeif.me/Go-fun/spider-training/infra/errors"
)

type PiaofangRankInfo struct {
	ID              int    `json:"-"`
	Rank            string `json:"rank"`
	MovieName       string `json:"movie_name"`
	ReleaseDate     string `json:"release_date"`
	Category        string `json:"category"`
	Director        string `json:"director"`
	BoxOfficeIncome string `json:"box_office_income"`
}

var (
	// <td class="num"><u>全球电影票房排行榜</u>第 1 名</td>
	rankPattern = `<tr><td class="num">(.*?)</td>`
	// <td class="title">《阿凡达》<span>Avatar</span></td>
	movieNamePattern = `<td class="title">(.*?)<span>(.*?)</span></td>`
	// <td class="year">2009</td>
	releaseDatePattern = `<td class="year">(.*?)</td>`
	// <td class="type">科幻/动作<span></span></td>
	categoryPattern = `<td class="type">(.*?)<span></span></td>`
	// <td class="daoyan">詹姆斯·卡梅隆</td>
	directorPattern = `<td class="daoyan">(.*?)</td>`
	// <td class="piaofang"><span>2,787,965,087</span>$</td>
	boxOfficeIncomePattern = `<td class="piaofang"><span>(.*?)</span>(.*?)</td>`
)

func GetPiaofangRankInfo(url string) ([]PiaofangRankInfo, error) {
	responseByte, err := downloader.GetHttpResponse(url, true)
	if err != nil {
		return nil, errors.ErrorPFangRankInfo
	}
	responseString := string([]byte(responseByte))

	var rankList []string
	reRank := regexp.MustCompile(rankPattern)
	for index, submatch := range reRank.FindAllStringSubmatch(responseString, -1) {
		if len(submatch[1]) > 34 {
			submatch[1] = submatch[1][34:len(submatch[1])]
		}
		fmt.Println(index, adapter.StringClear(submatch[1]))
		rankList = append(rankList, adapter.StringClear(submatch[1]))
	}

	var movieNameList []string
	removieName := regexp.MustCompile(movieNamePattern)
	for index, submatch := range removieName.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, submatch[1])
		movieNameList = append(movieNameList, submatch[1])
	}

	var releaseDateList []string
	rereleaseDate := regexp.MustCompile(releaseDatePattern)
	for index, submatch := range rereleaseDate.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, submatch[1])
		releaseDateList = append(releaseDateList, submatch[1])
	}

	var categoryList []string
	recategory := regexp.MustCompile(categoryPattern)
	for index, submatch := range recategory.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, submatch[1])
		categoryList = append(categoryList, submatch[1])
	}

	var directorList []string
	redirector := regexp.MustCompile(directorPattern)
	for index, submatch := range redirector.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, submatch[1])
		directorList = append(directorList, submatch[1])
	}

	var boxOfficeIncomeList []string
	reboxOfficeIncome := regexp.MustCompile(boxOfficeIncomePattern)
	for index, submatch := range reboxOfficeIncome.FindAllStringSubmatch(responseString, -1) {
		fmt.Println(index, submatch[1]+"$")
		boxOfficeIncomeList = append(boxOfficeIncomeList, submatch[1]+"$")
	}
	var piaofangRankInfoList []PiaofangRankInfo
	for index := range rankList {
		piaofangRankInfo := PiaofangRankInfo{
			ID:              index,
			Rank:            rankList[index],
			MovieName:       movieNameList[index],
			ReleaseDate:     releaseDateList[index],
			Category:        categoryList[index],
			Director:        directorList[index],
			BoxOfficeIncome: boxOfficeIncomeList[index],
		}
		piaofangRankInfoList = append(piaofangRankInfoList, piaofangRankInfo)
	}
	fmt.Println(piaofangRankInfoList)
	return piaofangRankInfoList, nil
}

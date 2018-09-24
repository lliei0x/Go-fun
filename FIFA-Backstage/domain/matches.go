package domain

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Matches(doc *goquery.Document) {
	doc.Find("div.fi-matchlist div.fi-mu-list a.fi-mu__link").Each(func(i int, selection *goquery.Selection) {
		matches := model.Match{}

		matches.Date = adapter.StringClear(selection.Find("div.fi-mu__info div.fi-mu__info__datetime").Text())
		matches.GroupName = adapter.StringClear(selection.Find("div.fi-mu__info div").Eq(2).Text())
		matches.Location = adapter.StringClear(selection.Find("div.fi-mu__info div.fi__info__location div.fi__info__stadium").Text())
		matches.CountryLeft = adapter.StringClear(selection.Find("div.fi-mu__m div.fi-t.fi-i--4.home div.fi-t__n span").Eq(0).Text())
		matches.CountryRight = adapter.StringClear(selection.Find("div.fi-mu__m div.fi-t.fi-i--4.away div.fi-t__n span").Eq(0).Text())
		matches.Score = adapter.StringClear(selection.Find("div.fi-s-wrap div.fi-s__score.fi-s__date-HHmm").Text())
		matchNumber := adapter.StringClear(selection.Find("div.fi-mu__info div.fi__info__matchnumber span").Eq(1).Text())
		matches.MatchNumber, _ = strconv.Atoi(matchNumber)

		// log.Println(matches)
		if err := initiator.POSTGRES.Create(&matches).Error; err != nil {
			log.Printf("Create Matches Error: %v", err)
		} //插入爬取到的信息
	})
}

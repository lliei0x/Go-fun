package domain

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Classic(doc *goquery.Document) {
	doc.Find("div.d3-o-media-object--vertical.fi-o-media-object__teaser").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		classic := model.Classic{}

		classicURL, _ := selection.Find("div.fi-module-timeline__comp figure a").Attr("href")
		if classicURL == "" {
			return false
		}
		classic.URL = adapter.StringClear(config.RootURL + classicURL)
		classicNameYear := adapter.StringClear(selection.Find("div.d3-o-media-object__body p").Text())
		classic.Name, classic.Year = strings.Split(classicNameYear, " ")[0], strings.Split(classicNameYear, " ")[1]

		worldCupInfo, err := download.Downloader(classic.URL)
		if err != nil {
			log.Println(err)
		}
		classicImage, _ := worldCupInfo.Find("div.ph-lnd-12 a img").Attr("src")
		classic.Image = adapter.StringClear("https" + classicImage)
		classic.Winner = adapter.StringClear(worldCupInfo.Find("div.c-team-rank div.t-n").Eq(0).Find("span.t-nText a").Text())
		classic.RunnersUp = adapter.StringClear(worldCupInfo.Find("div.c-team-rank div.t-n").Eq(1).Find("span.t-nText a").Text())
		classic.Third = adapter.StringClear(worldCupInfo.Find("div.c-team-rank div.t-n").Eq(2).Find("span.t-nText a").Text())
		classic.Fourth = adapter.StringClear(worldCupInfo.Find("div.c-team-rank div.t-n").Eq(3).Find("span.t-nText a").Text())
		classic.FinalResult = adapter.StringClear(worldCupInfo.Find("div.matches div.s-score.s-date-HHmm span").Text())
		classic.Title = adapter.StringClear(worldCupInfo.Find("h1.title a").Text())

		// log.Println(classic)
		if err := initiator.POSTGRES.Create(&classic).Error; err != nil {
			log.Printf("Get Classic err: %v", err)
		}

		return true
	})
}

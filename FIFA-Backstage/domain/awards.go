package domain

import (
	"log"

	"leeif.me/Go-fun/FIFA-Backstage/infra/config"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Awards(doc *goquery.Document) {
	doc.Find("section.fi-promo--enriched").Each(func(i int, selection *goquery.Selection) {
		award := model.Award{}
		// AwardName string `gorm:"type:varchar(64);not null;column:award_name"`
		// URL       string `gorm:"type:varchar(128);not null;column:url"`
		// Info      string `gorm:"type:varchar(128);not null;column:info"`
		award.AwardName = adapter.StringClear(selection.Find("h1").Text())
		awardURL, _ := selection.Find("div.fi-promo__image figure a").Attr("href")
		award.URL = adapter.StringClear(config.RootURL + awardURL)
		award.Info = adapter.StringClear(selection.Find("div.fi-promo-alignement-left div.fi-promo__summary p").Text())

		log.Println(award)
	})
}

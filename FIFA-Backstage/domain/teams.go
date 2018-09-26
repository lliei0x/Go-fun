package domain

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Teams(doc *goquery.Document) {
	doc.Find("div.fi-teams-list div.col-xs-6.col-sm-4.col-md-3.col-lg-3.col-flex a.fi-team-card.fi-team-card__team").Each(func(i int, selection *goquery.Selection) {
		teams := model.Team{}
		teamsflag, ok := selection.Find("div.fi-team-card__flag img").Attr("src")
		if !ok {
			log.Printf("Teams Images Error")
		}
		teams.Flag = adapter.StringClear(teamsflag)

		teams.Name = adapter.StringClear(selection.Find("div.fi-team-card__info div.fi-team-card__name").Text())

		// log.Println(teams)
		if err := initiator.POSTGRES.Create(&teams).Error; err != nil {
			log.Printf("Create Teams Error: %v", err)
		} //插入爬取到的信息
	})
}

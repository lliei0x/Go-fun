package domain

import (
	"log"
	"strconv"

	"leeif.me/Go-fun/FIFA-Backstage/infra/init"

	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Players(doc *goquery.Document) {
	doc.Find("div.fi-p--hub div a div.fi-p").Each(func(i int, selection *goquery.Selection) {
		players := model.Player{}
		players.Number, _ = strconv.Atoi(selection.Find("div.fi-p__jerseyNum  span").Text())
		players.Name = adapter.StringClear(selection.Find("div.fi-p__name").Text())
		players.Country = adapter.StringClear(selection.Find("div.fi-p__country").Text())
		players.Role = adapter.StringClear(selection.Find("div.fi-p__role").Text())
		playerImage, ok := selection.Find("div.fi-p__picture svg.fi-clip-svg image").Attr("href")
		if !ok {
			log.Println("Get player image err")
		}
		players.ImageURL = adapter.StringClear(playerImage)

		// log.Println(players)
		if err := initiator.POSTGRES.Create(&players).Error; err != nil {
			log.Printf("Create Players Error: %v", err)
		}
	})
}

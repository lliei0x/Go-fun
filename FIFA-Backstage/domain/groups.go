package domain

import (
	"log"
	"strconv"

	"leeif.me/Go-fun/FIFA-Backstage/infra/init"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func Groups(doc *goquery.Document) {
	groups := model.Group{}
	var (
		count     int
		groupList []string
	)
	doc.Find("div.fi-standings-list caption.fi-table__caption.clearfix").Each(func(i int, selection *goquery.Selection) {
		group := adapter.StringClear(selection.Find("p.fi-table__caption__title.fi-ltr--force").Text())
		groupList = append(groupList, group)
	})
	doc.Find("div.fi-standings-list tbody tr").Each(func(i int, selection *goquery.Selection) {
		if _, ok := selection.Attr("data-team-id"); ok != true {
			return
		}
		groups.GroupName = groupList[count/4]
		groups.TeamName = adapter.StringClear(selection.Find("div.fi-t.fi-i--4 div.fi-t__n span").Eq(0).Text())
		groups.MatchPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__matchplayed span").Text())
		groups.WinNumber, _ = strconv.Atoi(selection.Find("td.fi-table__win span").Text())
		groups.Draw, _ = strconv.Atoi(selection.Find("td.fi-table__draw span").Text())
		groups.Lost, _ = strconv.Atoi(selection.Find("td.fi-table__lost span").Text())
		groups.GoalFor, _ = strconv.Atoi(selection.Find("td.fi-table__goalfor span").Text())
		groups.GoalAgainst, _ = strconv.Atoi(selection.Find("td.fi-table__goalagainst span").Text())
		groups.DiffGoal, _ = strconv.Atoi(selection.Find("td.fi-table__diffgoal span").Text())
		groups.Points, _ = strconv.Atoi(selection.Find("td.fi-table__pts span").Text())
		count++

		// log.Println(groups)
		if err := initiator.POSTGRES.Create(&groups).Error; err != nil {
			log.Printf("Creste Groups Error: %v", err)
		}
	})
}

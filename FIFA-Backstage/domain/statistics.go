package domain

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func PlayerStatisticsWithGoalsScored(doc *goquery.Document) {
	doc.Find("div.col-sm-12 table tbody tr").Each(func(i int, selection *goquery.Selection) {
		playerGoalsScored := model.PlayersStatisticWithGoalsScored{}

		playerGoalsScored.Rank, _ = strconv.Atoi(selection.Find("td.fi-table__rank sorting_1 span").Text())
		playerGoalsScored.PlayerName = adapter.StringClear(selection.Find("td.fi-table__playername a div div.fi-p__wrapper-text div.fi-p__name").Text())
		playerGoalsScored.GoalsScored, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(0).Text())
		playerGoalsScored.Assists, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(1).Text())
		playerGoalsScored.MinutesPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(2).Text())
		playerGoalsScored.MatchesPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(3).Text())
		playerGoalsScored.PenaltiesScored, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(4).Text())
		playerGoalsScored.GoalsScoredLeft, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(5).Text())
		playerGoalsScored.GoalsScoredRight, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(6).Text())
		playerGoalsScored.HeadedGoals, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(7).Text())

		log.Println(playerGoalsScored)
	})
}

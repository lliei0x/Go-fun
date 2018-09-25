package domain

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"leeif.me/Go-fun/FIFA-Backstage/infra/adapter"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func PlayerStatisticsWithGoalsScored(doc *goquery.Document) {
	doc.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {
		playerGoalsScored := model.PlayersStatisticWithGoalsScored{}

		playerGoalsScored.Rank, _ = strconv.Atoi(selection.Find("td.fi-table__rank span").Text())
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

func PlayersStatisticWithShot(doc *goquery.Document) {
	doc.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {
		playersShot := model.PlayersStatisticWithShot{}

		playersShot.Rank, _ = strconv.Atoi(selection.Find("td.fi-table__rank span").Text())
		playersShot.Player = adapter.StringClear(selection.Find("td.fi-table__playername a div div.fi-p__wrapper-text div.fi-p__name").Text())
		playersShot.MatchesPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(0).Text())
		playersShot.MinutesPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(1).Text())
		playersShot.Shots, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(2).Text())
		playersShot.AttemptsOnTarget, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(3).Text())
		playersShot.AttemptsOffTarget, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(4).Text())
		playersShot.ShotsBlocked, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(5).Text())

		log.Println(playersShot)
	})
}

func TeamStatisticWithTopGoal(doc *goquery.Document) {
	doc.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {
		teamTopGoal := model.TeamStatisticWithTopGoal{}

		teamTopGoal.Rank, _ = strconv.Atoi(selection.Find("td.fi-table__rank span").Text())
		teamTopGoal.TeamName = adapter.StringClear(selection.Find("td.fi-table__teamname a div div.fi-t__n span").Eq(0).Text())
		teamTopGoal.MatchPlayed, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(0).Find("span").Text())
		teamTopGoal.GoalsFor, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(1).Find("span").Text())
		teamTopGoal.GoalsAgainst, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(2).Find("span").Text())
		teamTopGoal.PenaltyGoal, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(3).Find("span").Text())
		teamTopGoal.OwnGoals, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(4).Find("span").Text())
		teamTopGoal.OpenPlayGoals, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(5).Find("span").Text())
		teamTopGoal.SetPieceGoals, _ = strconv.Atoi(selection.Find("td.fi-table__td").Eq(6).Find("span").Text())

		log.Println(teamTopGoal)
	})
}

package main

import (
	"leeif.me/Go-fun/FIFA-Backstage/domain"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func main() {
	// create table by gorm auto migrate
	CreateTable()
	defer initiator.POSTGRES.Close()

	// push data into db
	PushDataToDb()

	// start http server
}

// CreateTable create table by gorm auto migrate
func CreateTable() {
	initiator.POSTGRES.AutoMigrate(&model.Match{},
		&model.Award{},
		&model.Group{},
		&model.Team{},
		&model.Player{},
		&model.Classic{},
		&model.TeamStatisticWithTopGoal{},
		&model.PlayersStatisticWithGoalsScored{},
		&model.PlayersStatisticWithShot{},
		// &model.Admin{},
	)
}

func PushDataToDb() {
	doc, _ := download.Downloader(config.MatchesURL)
	domain.Matches(doc)

	doc, _ = download.Downloader(config.AwardsURL)
	domain.Awards(doc)

	doc, _ = download.Downloader(config.GroupsURL)
	domain.Groups(doc)

	doc, _ = download.Downloader(config.TeamsURL)
	domain.Teams(doc)

	doc, _ = download.Downloader(config.PlayersURLList)
	domain.Players(doc)

	doc, _ = download.Downloader(config.ClassicURL)
	domain.Classic(doc)

	doc, _ = download.Downloader(config.StatisticsTeamGoalURL)
	domain.TeamStatisticWithTopGoal(doc)

	doc, _ = download.Downloader(config.StatisticsPlayerGoalScoredURL)
	domain.PlayerStatisticsWithGoalsScored(doc)

	doc, _ = download.Downloader(config.StatisticsPlayerShots)
	domain.PlayersStatisticWithShot(doc)

}

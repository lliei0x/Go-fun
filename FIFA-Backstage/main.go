package main

import (
	_ "leeif.me/Go-fun/FIFA-Backstage/docs"
	"leeif.me/Go-fun/FIFA-Backstage/domain"
	"leeif.me/Go-fun/FIFA-Backstage/infra/config"
	"leeif.me/Go-fun/FIFA-Backstage/infra/download"
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
	"leeif.me/Go-fun/FIFA-Backstage/ui/api-server"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host

func main() {
	// create table by gorm auto migrate
	CreateTable()
	defer initiator.POSTGRES.Close()

	// push data into db
	// PushDataToDb()

	// start http server
	api_server.New().Start()
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

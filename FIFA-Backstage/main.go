package main

import (
	"leeif.me/Go-fun/FIFA-Backstage/infra/init"
	"leeif.me/Go-fun/FIFA-Backstage/infra/model"
)

func main() {
	// create table by gorm auto migrate
	// CreateTable()
	// defer initiator.POSTGRES.Close()

	// push data into db
	PushDataToDb()

	// start http server
}

// CreateTable create table by gorm auto migrate
func CreateTable() {
	initiator.POSTGRES.AutoMigrate(&model.Match{},
		&model.Award{},
		&model.Coach{},
		&model.Group{},
		&model.Team{},
		&model.Player{},
		&model.WorldCupArchive{},
		&model.TeamStatisticWithTopGoal{},
		&model.PlayersStatisticWithGoalsScored{},
		&model.PlayersStatisticWithShot{},
		&model.Admin{},
	)
}

func PushDataToDb() {
	// matches
	// docs, err := download.Downloader(config.MatchesURL)
}

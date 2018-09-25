package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	MatchesURL                    = "https://www.fifa.com/worldcup/matches"
	TeamsURL                      = "https://www.fifa.com/worldcup/teams/"
	GroupsURL                     = "https://www.fifa.com/worldcup/groups/"
	PlayersURL                    = "https://www.fifa.com/worldcup/players/browser/#player-by-position"
	PlayersURLList                = "https://www.fifa.com/worldcup/players/_libraries/byposition/all/_players-list"
	StatisticsURL                 = "https://www.fifa.com/worldcup/statistics/"
	StatisticsPlayerGoalScoredURL = "https://www.fifa.com/worldcup/statistics/players/goal-scored"
	StatisticsPlayerShots         = "https://www.fifa.com/worldcup/statistics/players/shots"
	StatisticsTeamGoalURL         = "https://www.fifa.com/worldcup/statistics/teams/goal-scored"
	StatisticsTeamShots           = "https://www.fifa.com/worldcup/statistics/teams/shots"
	AwardsURL                     = "https://www.fifa.com/worldcup/awards/"
	ClassicURL                    = "https://www.fifa.com/worldcup/classic/"
)

var (
	RootURL = "https://www.fifa.com"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../infra/config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetPostgreConfg() string {
	var (
		dbName   string
		port     string
		user     string
		sslMode  string
		password string
		host     string
	)

	dbName = viper.GetString("db.dbname")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	sslMode = viper.GetString("db.sslmode")
	password = viper.GetString("db.password")
	host = viper.GetString("db.host")

	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password)
}

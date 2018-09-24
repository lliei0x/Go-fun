package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PlayersStatisticWithGoalsScored struct {
	gorm.Model
	Rank             int    `gorm:"type:integer; not null; column:rank"`
	PlayerName       string `gorm:"type:varchar(32); not null; column:player_name"`
	GoalsScored      int    `gorm:"type:integer; not null; column:goals_scored"`
	Assists          int    `gorm:"type:integer; not null; column:assists"`
	MinutesPlayed    int    `gorm:"type:integer; not null; column:minutes_played"`
	MatchesPlayed    int    `gorm:"type:integer; not null; column:matches_played"`
	PenaltiesScored  int    `gorm:"type:integer; not null; column:penalties_scored"`
	GoalsScoredLeft  int    `gorm:"type:integer; default:0; column:goals_scored_left"`
	GoalsScoredRight int    `gorm:"type:integer; default:0; column:goals_scored_right"`
	HeadedGoals      int    `gorm:"type:integer; default:0; column:headed_goals"`
}

type PlayersStatisticWithGoalsScoredSerializer struct {
	ID               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdateAt         time.Time  `json:"update_at"`
	DeleteAt         *time.Time `json:"delete_at"`
	Rank             int        `json:"rank"`
	PlayerName       string     `json:"player_name"`
	GoalsScored      int        `json:"goals_scored"`
	Assists          int        `json:"assists"`
	MinutesPlayed    int        `json:"minutes_played"`
	MatchesPlayed    int        `json:"matches_played"`
	PenaltiesScored  int        `json:"penalties_scored"`
	GoalsScoredLeft  int        `json:"goals_scored_left"`
	GoalsScoredRight int        `json:"goals_scored_right"`
	HeadedGoals      int        `json:"headed_goals"`
}

func (pss *PlayersStatisticWithGoalsScored) Serializer() PlayersStatisticWithGoalsScoredSerializer {
	return PlayersStatisticWithGoalsScoredSerializer{
		ID:               pss.ID,
		CreatedAt:        pss.CreatedAt,
		UpdateAt:         pss.UpdatedAt,
		DeleteAt:         pss.DeletedAt,
		Rank:             pss.Rank,
		PlayerName:       pss.PlayerName,
		GoalsScored:      pss.GoalsScored,
		Assists:          pss.Assists,
		MinutesPlayed:    pss.MinutesPlayed,
		MatchesPlayed:    pss.MatchesPlayed,
		PenaltiesScored:  pss.PenaltiesScored,
		GoalsScoredLeft:  pss.GoalsScoredLeft,
		GoalsScoredRight: pss.GoalsScoredRight,
		HeadedGoals:      pss.HeadedGoals,
	}
}

type PlayersStatisticWithShot struct {
	gorm.Model
	Rank              int    `gorm:"type:integer; not null; column:rank"`
	Player            string `gorm:"type:varchar(32); not null; column:player_name"`
	MatchesPlayed     int    `gorm:"type:integer;column:matches_played"`
	MinutesPlayed     int    `gorm:"type:integer;column:minutes_played"`
	Shots             int    `gorm:"type:integer;column:shots_number"`
	AttemptsOnTarget  int    `gorm:"type:integer;column:attempts_On_target"`
	AttemptsOffTarget int    `gorm:"type:integer;column:attempts_off_target"`
	ShotsBlocked      int    `gorm:"type:integer;column:shots_blocked"`
}

type PlayersStatisticWithShotSerializer struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdateAt          time.Time  `json:"update_at"`
	DeleteAt          *time.Time `json:"delete_at"`
	Rank              int        `json:"rank"`
	Player            string     `json:"player"`
	MatchesPlayed     int        `json:"matches_played"`
	MinutesPlayed     int        `json:"minutes_played"`
	Shots             int        `json:"shots"`
	AttemptsOnTarget  int        `json:"attempts_on_target"`
	AttemptsOffTarget int        `json:"attempts_off_target"`
	ShotsBlocked      int        `json:"shots_blocked"`
}

func (ps *PlayersStatisticWithShot) Serializer() PlayersStatisticWithShotSerializer {
	return PlayersStatisticWithShotSerializer{
		ID:                ps.ID,
		CreatedAt:         ps.CreatedAt,
		UpdateAt:          ps.UpdatedAt,
		DeleteAt:          ps.DeletedAt,
		Rank:              ps.Rank,
		Player:            ps.Player,
		MatchesPlayed:     ps.MatchesPlayed,
		MinutesPlayed:     ps.MinutesPlayed,
		Shots:             ps.Shots,
		AttemptsOnTarget:  ps.AttemptsOnTarget,
		AttemptsOffTarget: ps.AttemptsOffTarget,
		ShotsBlocked:      ps.ShotsBlocked,
	}
}

type TeamStatisticWithTopGoal struct {
	gorm.Model
	Rank          int    `gorm:"type:integer; not null; column:rank"`
	TeamName      string `gorm:"type:varchar(12); not null; column:team_name"`
	MatchPlayed   int    `gorm:"type:integer; not null; column:matches_payed"`
	GoalsFor      int    `gorm:"type:integer; not null; column:goals_for"`
	GoalsScored   int    `gorm:"type:integer; not null; column:goals_scored"`
	GoalsAgainst  int    `gorm:"type:integer; not null; column:goals_against"`
	PenaltyGoal   int    `gorm:"type:integer; not null; column:penalty_goal"`
	OwnGoals      int    `gorm:"type:integer; not null; column:own_goals"`
	OpenPlayGoals int    `gorm:"type:integer; not null; column:open_play_goals"`
	SetPieceGoals int    `gorm:"type:integer; not null; column:set_piece_goals"`
}

type TeamStatisticWithTopGoalSerializer struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdateAt      time.Time  `json:"update_at"`
	DeleteAt      *time.Time `json:"delete_at"`
	Rank          int        `json:"int"`
	TeamName      string     `json:"team_name"`
	MatchPlayed   int        `json:"match_played"`
	GoalFor       int        `json:"goal_for"`
	GoalsScored   int        `json:"goals_scored"`
	GoalsAgainst  int        `json:"goals_against"`
	PenaltyGoal   int        `json:"penalty_goal"`
	OwnGoals      int        `json:"own_goals"`
	OpenPlayGoals int        `json:"open_play_goals"`
	SetPieceGoals int        `json:"set_piece_goals"`
}

func (ttg *TeamStatisticWithTopGoal) Serializer() TeamStatisticWithTopGoalSerializer {
	return TeamStatisticWithTopGoalSerializer{
		ID:            ttg.ID,
		CreatedAt:     ttg.CreatedAt,
		UpdateAt:      ttg.UpdatedAt,
		DeleteAt:      ttg.DeletedAt,
		Rank:          ttg.Rank,
		TeamName:      ttg.TeamName,
		MatchPlayed:   ttg.MatchPlayed,
		GoalFor:       ttg.GoalsFor,
		GoalsScored:   ttg.GoalsScored,
		GoalsAgainst:  ttg.GoalsAgainst,
		PenaltyGoal:   ttg.PenaltyGoal,
		OwnGoals:      ttg.OwnGoals,
		OpenPlayGoals: ttg.OpenPlayGoals,
		SetPieceGoals: ttg.SetPieceGoals,
	}
}

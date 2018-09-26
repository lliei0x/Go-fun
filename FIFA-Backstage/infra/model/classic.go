package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Classic struct {
	gorm.Model
	URL         string `gorm:"type:varchar(128);column:world_cup_url"`
	Name        string `gorm:"type:varchar(128);column:country_name"`
	Year        string `gorm:"type:varchar(64);column:year"`
	Image       string `gorm:"type:varchar(128);column:image"`
	Winner      string `gorm:"type:varchar(64);column:winner_country"`
	RunnersUp   string `gorm:"type:varChar(64);column:runners_up_name"`
	Third       string `gorm:"type:varchar(64);column:third_name"`
	Fourth      string `gorm:"type:varchar(64);column:fourth_name"`
	FinalResult string `gorm:"type:varchar(128);column:final_result"`
	Title       string `gorm:"type:varchar(128);column:title"`
}

type ClassicSerializer struct {
	ID          uint       `json:"id"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
	DeleteAt    *time.Time `json:"delete_at"`
	URL         string     `json:"url"`
	Name        string     `json:"name"`
	Year        string     `json:"year"`
	Image       string     `json:"image"`
	Winner      string     `json:"winner"`
	RunnersUp   string     `json:"runners_up"`
	Third       string     `json:"third_name"`
	Fourth      string     `json:"fourth_name"`
	FinalResult string     `json:"final_result"`
	Title       string     `json:"title"`
}

func (c *Classic) Serializer() ClassicSerializer {
	return ClassicSerializer{
		ID:          c.ID,
		CreateAt:    c.CreatedAt,
		UpdateAt:    c.UpdatedAt,
		DeleteAt:    c.DeletedAt,
		URL:         c.URL,
		Name:        c.Name,
		Year:        c.Year,
		Image:       c.Image,
		Winner:      c.Winner,
		RunnersUp:   c.RunnersUp,
		Third:       c.Third,
		Fourth:      c.Fourth,
		FinalResult: c.FinalResult,
		Title:       c.Title,
	}
}

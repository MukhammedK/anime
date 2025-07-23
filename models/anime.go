package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Anime struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Studio      string
	Type        string
	Episodes    string
	Author      string
	Genres      pq.StringArray `gorm:"type:text[]"`
	Rating      float64
	Description string
	CoverUrl    string `gorm:"column=cover_url"`
	ReleaseYear int    `gorm:"column=release_year"`
}

func (Anime) TableName() string {
	return "anime"
}

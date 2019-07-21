package model

import "github.com/jinzhu/gorm"

type Match struct {
	gorm.Model
	Date         string `gorm:"type:varchar(64);not null;column:date"`
	GroupName    string `gorm:"type:varchar(64);not null;column:group_name"`
	Location     string `gorm:"type:varchar(64);not null;column:location"`
	CountryLeft  string `gorm:"type:varchar(64);not null;column:country_left"`
	CountryRight string `gorm:"type:varchar(64);not null;column:country_right"`
	Score        string `gorm:"type:varchar(64);not null;column:score"`
	MatchNumber  int    `gorm:"type:integer;not null;column:match_number"`
}

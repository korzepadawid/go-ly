package model

import (
	"github.com/korzepadawid/go-ly/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type URL struct {
	gorm.Model
	Url        string
	AuthorIPv4 string
}

func InitURL() {
	db = config.GetDB()
	db.AutoMigrate(&URL{})
}

type Shortener interface {
	SaveURL() *URL
	GetURLByID(ID int64) *URL
}

func (u *URL) SaveURL() *URL {
	db.Create(&u)
	return u
}

func GetURLByID(ID int64) *URL {
	var url URL
	db.Where("ID=?", ID).Find(&url)
	return &url
}

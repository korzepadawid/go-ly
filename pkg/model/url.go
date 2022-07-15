package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Url        string
	AuthorIPv4 string
}

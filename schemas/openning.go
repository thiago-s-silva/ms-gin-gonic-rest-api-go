package schemas

import (
	"gorm.io/gorm"
)

type Openning struct {
	gorm.Model // unrap
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     int64
}

package models

import (
	"time"

	"github.com/rohanshrestha09/todo/enums"
)

type Todo struct {
	Model
	Name      string           `json:"name" gorm:"not null"`
	Start     time.Time        `json:"start" gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Due       time.Time        `json:"due" gorm:"not null"`
	Note      string           `json:"note" gorm:"type:text"`
	File      string           `json:"file"`
	FileName  string           `json:"filename"`
	Status    enums.TodoStatus `json:"status" gorm:"type:enum('IN_PROGRESS','COMPLETED');default:IN_PROGRESS;not null"`
	Important bool             `json:"important" gorm:"default:false;not null"`
	UserID    uint             `json:"userId" gorm:"not null"`
	User      *User            `json:"user"`
	ListID    uint             `json:"listId" gorm:"not null"`
	List      *List            `json:"list"`
}

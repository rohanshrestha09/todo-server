package models

type List struct {
	Model
	Name   string  `json:"name" gorm:"not null"`
	UserID uint    `json:"userId" gorm:"not null"`
	User   *User   `json:"user"`
	Todos  []*Todo `json:"todos"`
}

package models

type User struct {
	Model
	Name      string  `json:"name" gorm:"not null"`
	Username  string  `json:"username" gorm:"not null;unique"`
	Email     string  `json:"email,omitempty" gorm:"not null;unique"`
	Password  string  `json:"-" gorm:"not null"`
	Image     string  `json:"image"`
	ImageName string  `json:"imagename"`
	Provider  string  `json:"provider" gorm:"type:enum('GOOGLE','EMAIL');default:EMAIL;not null"`
	Lists     []*List `json:"lists"`
	Todos     []*Todo `json:"todos"`
}

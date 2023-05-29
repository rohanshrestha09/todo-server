package dto

import (
	"mime/multipart"
	"time"
)

type RegisterDTO struct {
	Name            string `form:"name" binding:"required"`
	Username        string `form:"username" binding:"required"`
	Email           string `form:"email" binding:"required,email" validate:"email"`
	Password        string `form:"password" binding:"required" validate:"gte=8"`
	ConfirmPassword string `form:"confirmPassword" binding:"required" validate:"eqfield=Password"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileUpdateDTO struct {
	Name  string                `form:"name"`
	Image *multipart.FileHeader `form:"image"`
}

type ListCreateDTO struct {
	Name string `json:"name" binding:"required"`
}

type ListUpdateDTO struct {
	Name string `json:"name"`
}

type TodoCreateDTO struct {
	Name      string                `form:"name" binding:"required"`
	Due       time.Time             `form:"due" binding:"required"`
	Note      string                `form:"note"`
	File      *multipart.FileHeader `form:"file"`
	Important bool                  `form:"important"`
}

type TodoUpdateDTO struct {
	Name string                `form:"name"`
	Due  time.Time             `form:"due"`
	Note string                `form:"note"`
	File *multipart.FileHeader `form:"file"`
}

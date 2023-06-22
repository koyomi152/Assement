package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true"form:"id"json:"id"`
	Username string `gorm:"unique;not null"form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null"form:"password" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null"form:"email"json:"email"binding:"required"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

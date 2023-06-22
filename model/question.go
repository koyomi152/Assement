package model

type Question struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true"form:"id"json:"id"`
	Title    string `gorm:"not null"form:"title"json:"title"`
	Content  string `gorm:"not null"form:"content"json:"content"`
	Username string `gorm:"not null"form:"username"json:"username"`
}

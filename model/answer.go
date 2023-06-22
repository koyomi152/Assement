package model

type Answer struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true"form:"id"json:"id"`
	QuestionID uint   `gorm:"index"`
	Content    string `gorm:"not null"form:"content"json:"content"`
	Username   string `gorm:"not null"form:"username"json:"username"`
}

package dao

import (
	"ginHomework/model"
	"gorm.io/gorm"
)

type QuestionDao struct {
	DB *gorm.DB
}

func NewQuestionDao(db *gorm.DB) *QuestionDao {
	return &QuestionDao{
		DB: db,
	}
}
func (dao *QuestionDao) CreateQuestion(question *model.Question) error {
	return dao.DB.Create(question).Error
}

func (dao *QuestionDao) GetQuestionsByUsername(username string) ([]*model.Question, error) {
	var questions []*model.Question
	err := dao.DB.Where("username = ?", username).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
func (dao *QuestionDao) GetQuestionsByTitle(title string) ([]*model.Question, error) {
	var questions []*model.Question
	err := dao.DB.Where("title = ?", title).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
func (dao *QuestionDao) DeleteQuestionByQuestionID(questionID uint) error {
	return dao.DB.Where("id = ?", questionID).Delete(&model.Question{}).Error
}

func (dao *QuestionDao) UpdateQuestionContentByUsername(username string, questionID int64, content string) error {
	return dao.DB.Model(&model.Question{}).
		Where("username = ? AND id = ?", username, questionID).
		Update("content", content).
		Error
}

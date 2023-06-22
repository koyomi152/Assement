package dao

import (
	"ginHomework/model"
	"gorm.io/gorm"
)

type AnswerDao struct {
	DB *gorm.DB
}

func NewAnswerDao(db *gorm.DB) *AnswerDao {
	return &AnswerDao{
		DB: db,
	}
}

func (dao *AnswerDao) CreateAnswer(answer *model.Answer) error {
	return dao.DB.Create(answer).Error
}
func (dao *AnswerDao) GetAnswersByUsername(username string) ([]*model.Answer, error) {
	var answers []*model.Answer
	err := dao.DB.Where("username = ?", username).Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}
func (dao *AnswerDao) DeleteAnswersByQuestionID(questionID uint) error {
	return dao.DB.Where("question_id = ?", questionID).Delete(&model.Answer{}).Error
}
func (dao *AnswerDao) DeleteAnswerByAnswerID(answerID int64) error {
	return dao.DB.Where("id = ?", answerID).Delete(&model.Answer{}).Error
}
func (dao *AnswerDao) UpdateAnswerContentByUsername(username string, answerID int64, content string) error {
	return dao.DB.Model(&model.Answer{}).
		Where("username = ? AND id = ?", username, answerID).
		Update("content", content).
		Error
}

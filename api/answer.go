package api

import (
	"ginHomework/dao"
	"ginHomework/model"
	"ginHomework/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func answerQuestion(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		utils.RespFail(c, "验证信息错误，请重新验证")
		return
	}

	// 将值转换为字符串类型
	answerUsername, ok := username.(string)
	if !ok {
		return
	}
	// 从请求中获取问题ID和回答内容
	questionIDStr := c.Param("id")
	answerContent := c.PostForm("content")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	var answer model.Answer
	answer.QuestionID = uint(questionID)
	answer.Content = answerContent
	answer.Username = answerUsername
	answerdao := dao.NewAnswerDao(db)
	answerdao.CreateAnswer(&answer)
}
func deleteAnswer(c *gin.Context) {
	answerIDStr := c.Param("answerID")

	// 转换类型
	answerID, err := strconv.ParseUint(answerIDStr, 10, 64)
	if err != nil {
		utils.RespFail(c, "invalid answer ID")
		return
	}

	answerDao := dao.NewAnswerDao(db)

	// 删除 Answer
	err = answerDao.DeleteAnswerByAnswerID(int64(answerID))
	if err != nil {
		utils.RespFail(c, "failed to delete answer")
		return
	}

	// 成功删除 Answer
	utils.RespSuccess(c, "Answer deleted successfully")
}
func updateAnswerContent(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		utils.RespFail(c, "验证信息错误，请重新验证")
		return
	}
	// 将值转换为字符串类型
	answerUsername, ok := username.(string)
	if !ok {
		return
	}

	// 从请求中获取回答ID和新的内容
	answerIDStr := c.Param("id")
	answerContent := c.PostForm("content")

	// 调用 DAO 层函数更新 Answer 内容
	answerID, err := strconv.ParseUint(answerIDStr, 10, 64)
	if err != nil {
		utils.RespFail(c, "invalid answer ID")
		return
	}

	answerDao := dao.NewAnswerDao(db)
	err = answerDao.UpdateAnswerContentByUsername(answerUsername, int64(answerID), answerContent)
	if err != nil {
		utils.RespFail(c, "更新回答内容失败")
		return
	}

	// 成功更新 Answer 内容
	utils.RespSuccess(c, "回答内容更新成功")
}

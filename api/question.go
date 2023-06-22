package api

import (
	"ginHomework/dao"
	"ginHomework/model"
	"ginHomework/utils"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"strconv"
)

func createQuestion(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		utils.RespFail(c, "验证信息错误，请重新验证")
		return
	}

	// 将值转换为字符串类型
	questionUsername, ok := username.(string)
	if !ok {
		return
	}
	var question model.Question
	question.Title = c.PostForm("title")
	question.Content = c.PostForm("content")
	question.Username = questionUsername
	questiondao := dao.NewQuestionDao(db)
	questiondao.CreateQuestion(&question)

}
func getQuestions(c *gin.Context) {
	questiondao := dao.NewQuestionDao(db)

	// 获取请求参数
	title := c.Query("title")

	// 根据标题查找问题内容
	questions, err := questiondao.GetQuestionsByTitle(title)
	if err != nil {
		utils.RespFail(c, "failed to get questions")
		return
	}
	// 返回查询结果
	jsonData, err := json.Marshal(questions)
	if err != nil {
		utils.RespFail(c, "failed to serialize questions")
		return
	}

	// 返回查询结果
	utils.RespSuccess(c, string(jsonData))
}
func deleteQuestion(c *gin.Context) {
	questionIDStr := c.Param("id")
	//转换类型
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		utils.RespFail(c, "invalid question ID")
		return
	}
	answerDao := dao.NewAnswerDao(db)
	answerDao.DeleteAnswersByQuestionID(uint(questionID))
	questionDao := dao.NewQuestionDao(db)

	// 删除 Question
	err = questionDao.DeleteQuestionByQuestionID(uint(questionID))
	if err != nil {
		utils.RespFail(c, "failed to delete question")
		return
	}

	// 成功删除 Question
	utils.RespSuccess(c, "Question deleted successfully")
}
func updateQuestionContent(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		utils.RespFail(c, "验证信息错误，请重新验证")
		return
	}

	// 将值转换为字符串类型
	questionUsername, ok := username.(string)
	if !ok {
		return
	}

	// 从请求中获取问题ID和新的内容
	questionIDStr := c.Param("id")
	questionContent := c.PostForm("content")

	// 调用 DAO 层函数更新 Question 内容
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "无效的问题ID")
		return
	}

	questionDao := dao.NewQuestionDao(db)
	err = questionDao.UpdateQuestionContentByUsername(questionUsername, int64(questionID), questionContent)
	if err != nil {
		utils.RespFail(c, "更新问题内容失败")
		return
	}

	// 成功更新 Question 内容
	utils.RespSuccess(c, "问题内容更新成功")
}

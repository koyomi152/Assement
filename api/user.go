package api

import (
	"fmt"
	"ginHomework/api/middleware"
	"ginHomework/dao"
	"ginHomework/model"
	"ginHomework/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB

func register(c *gin.Context) {
	//验证请求数据是否符合预期的格式和要求
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}
	var user model.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")

	userDao := dao.NewUserDao(db)
	// 验证用户名是否存在
	_, err := userDao.GetUserByUsername(user.Username)
	if err != nil {
		userDao.CreateUser(&user)
		// 以 JSON 格式返回信息
		utils.RespSuccess(c, "add user successful")
	} else {
		utils.RespFail(c, "user exists")
		return
	}

}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	userDao := dao.NewUserDao(db)
	var user model.User
	// 传入用户名和密码
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")

	// 验证用户名是否存在
	_, err := userDao.GetUserByUsername(user.Username)
	if err != nil {
		utils.RespFail(c, "user doesn't exists")
		return
	}
	// 查找正确的密码
	selectPassword, _ := userDao.SelectPasswordFromUsername(user.Username)
	// 若不正确则传出错误
	if selectPassword != user.Password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}
	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: user.Username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "L",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
func getUserInfo(c *gin.Context) {
	// 从URL参数中获取用户ID
	userID := c.Param("id")

	// 先从缓存中获取用户信息
	cacheKey := fmt.Sprintf("user:%s", userID)
	cachedData, err := redisClient.Get(redisClient.Context(), cacheKey).Result()
	if err == nil {
		// 如果缓存中存在用户信息，则直接返回缓存数据
		c.JSON(http.StatusOK, gin.H{
			"message": "User information retrieved from cache",
			"data":    cachedData,
		})
		return
	}

	// 模拟从数据库中获取用户信息
	userInfo := getUserInfoFromDB(userID)

	// 将用户信息存储到缓存中，并设置过期时间
	err = redisClient.Set(redisClient.Context(), cacheKey, userInfo, 5*time.Minute).Err()
	if err != nil {
		log.Println("Failed to set user info in cache:", err)
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"message": "User information retrieved from database",
		"data":    userInfo,
	})
}

// 模拟从数据库中获取用户信息
func getUserInfoFromDB(userID string) string {
	// 这里可以根据实际情况从数据库中查询用户信息
	// 这里只是简单地返回模拟数据
	return fmt.Sprintf("User ID: %s, Name: John Doe, Age: 30", userID)
}

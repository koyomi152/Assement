package api

import (
	"ginHomework/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	redisClient *redis.Client
)

func InitRedis() {
	// 初始化Redis客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // Redis服务器密码，如果有的话
		DB:       0,                // Redis数据库编号
	})

	// 检查Redis连接是否正常
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	// 定义获取用户信息的路由
	r.GET("/user/:id", getUserInfo)
	r.POST("/register", register) // 注册
	r.POST("/login", login)       // 登录
	r.GET("/questions", getQuestions)
	r.POST("/createQuestions", createQuestion)
	r.POST("/questions/:id/answers", answerQuestion)
	r.GET("/user/:id/questions", getQuestions)     // 获取用户的所有问题
	r.DELETE("/questions/:id", deleteQuestion)     // 删除问题
	r.PUT("/questions/:id", updateQuestionContent) // 修改问题
	r.DELETE("/answers/:id", deleteAnswer)         // 删除回答
	r.PUT("/answers/:id", updateAnswerContent)     // 修改回答

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8088") // 跑在 8088 端口上
}

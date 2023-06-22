package dao

import (
	"ginHomework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type UserDao struct {
	DB *gorm.DB
}

func InitDB() {
	dsn := "root:password@tcp(localhost:3306)/gindemo?charset=utf8mb4&parseTime=True&loc=Local"
	// 数据库连接信息，包括用户名、密码、数据库地址和端口，以及数据库名称
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	log.Println("connect success:")
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate user table: " + err.Error())
	}
}
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		DB: db,
	}
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(user).Error
}

// GetUserByUsername 根据用户名查询用户
func (dao *UserDao) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := dao.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (dao *UserDao) SelectPasswordFromUsername(username string) (string, error) {
	user := &model.User{}
	err := dao.DB.Select("password").Where("username = ?", username).First(user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

// UpdateUser 更新用户信息
func (dao *UserDao) UpdateUser(user *model.User) error {
	return dao.DB.Save(user).Error
}

// DeleteUser 删除用户
func (dao *UserDao) DeleteUser(user *model.User) error {
	return dao.DB.Delete(user).Error
}

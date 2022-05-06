package opdb

import (
	"errors"
	"jwtDemo/dfst"
	"time"
)

type User struct {
	Id    int32  `gorm:"AUTO_INCREMENT,primary_key"`
	Name  string `gorm:"name"`
	Pwd   string `gorm:"pwd"`
	Phone string `gorm:"phone"`
	Email string `gorm:"email"`
	//other extend
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InitModel() {
	DB.AutoMigrate(&User{}) //如果表不存在,通过User这个结构体建表
}

//查看该用户是否在数据库中
func UserIsExists(username string) bool {
	result := false
	// 指定库
	var user User
	dbResult := DB.Where("name = ?", username).Find(&user)
	if dbResult.Error != nil {
		result = false
	} else {
		result = true
	}
	return result
}

//将用户信息插入数据库中
func (user *User) Insert() error {
	return DB.Model(&User{}).Create(&user).Error
}
func Register(userInfo dfst.RegisterInfo) error {
	//查看该用户是否在数据库中
	if UserIsExists(userInfo.Name) {
		return errors.New("用户已经存在,请直接登陆")
	}
	user := User{
		Name:  userInfo.Name,
		Pwd:   userInfo.Pwd,
		Phone: userInfo.Phone,
		Email: userInfo.Email,
	}
	//将用户信息插入数据库中
	err := user.Insert()
	return err
}

//验证账号密码是否正确，即是否在数据库中存在，注册过
func LoginPass(login dfst.LoginReq) (pass bool, err error) {
	var user User
	dbErr := DB.Where("name = ? && pwd = ?", login.Name, login.Pwd).Find(&user).Error
	if dbErr != nil {
		return false, dbErr
	}
	return true, nil
}

package dfst

import "time"

//将结构体抽离出来，结果包循环引用的问题

type RegisterInfo struct {
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}

type LoginResult struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}

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

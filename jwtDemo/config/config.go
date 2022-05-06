package config

import (
	"gopkg.in/ini.v1"
	"jwtDemo/dfst"
	"log"
)

//解析ini文件并反射到结构体中
func parserConfig(dbCfg *dfst.DBConfig) {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Panicf("Fail to read file: %v\n", err)
	}
	dbCfg.Host = cfg.Section("DB").Key("host").String()
	dbCfg.Port, _ = cfg.Section("DB").Key("port").Int()
	dbCfg.User = cfg.Section("DB").Key("user").String()
	dbCfg.Pwd = cfg.Section("DB").Key("pwd").String()
	dbCfg.Database = cfg.Section("DB").Key("database").String()
}

//读取配置文件内容
func LoadDBConfig() *dfst.DBConfig {
	dbCfg := &dfst.DBConfig{}
	//解析ini文件并反射到结构体中
	parserConfig(dbCfg)
	return dbCfg
}

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Username string
	Password string
	Ip       string
	Port     string
	Database string
}

func main() {
	var mysqlcfg MysqlConfig
	if _, err := toml.DecodeFile("config/mysql.toml", &mysqlcfg); err != nil {
		panic(err)
	}
	dsn := mysqlcfg.Username + ":" + mysqlcfg.Password + "@tcp(" + mysqlcfg.Ip + ":" + mysqlcfg.Port + ")/" + mysqlcfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("db=%v, err=%v\n", db, err)
}

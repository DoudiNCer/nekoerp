package orm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *sql.DB
var Db *gorm.DB

type MysqlConfig struct {
	Host     string `json:"mySQLHost"`     // MySQL主机
	Port     int    `json:"mySQLPort"`     // MySQL端口
	DataBase string `json:"mySQLDataBase"` // MySQL数据库
	Username string `json:"mySQLUsername"` // MySQL用户名
	Password string `json:"mySQLPassword"` // MySQL密码
	IsDebug  bool   `json:"isDebug"`       // 调试模式
}

func NewOrm(config *MysqlConfig, orm *gorm.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DataBase)
	db, err := gorm.Open(mysql.Open(dsn), orm)
	if err != nil {
		return err
	}
	if config.IsDebug {
		db.Debug()
	}
	conn, err := db.DB()
	if err != nil {
		return err
	}
	Conn = conn
	Db = db
	return nil
}

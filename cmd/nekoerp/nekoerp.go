package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nekoerp/cmd/nekoerp/config"
	"nekoerp/cmd/nekoerp/service"
	orm "nekoerp/pkg/storage/mysql"
)

func main() {
	// 加载配置文件
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	// 创建 ORM
	err = orm.NewOrm(&orm.MysqlConfig{
		Host:     conf.MySQLHost,
		Port:     conf.MySQLPort,
		Username: conf.MySQLUsername,
		Password: conf.MySQLPassword,
		DataBase: conf.MySQLDataBase,
		IsDebug:  conf.IsDebug,
	}, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	defer orm.Conn.Close()

	// 初始化 gin
	if !conf.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	gin.Logger()
	service.RegisterHandler(app)
	err = app.Run(fmt.Sprintf("%s:%d", conf.WebHost, conf.WebPort))
	if err != nil {
		panic(err)
	}
}

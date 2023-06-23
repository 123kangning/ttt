package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"qin/configs/consts"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,                                //禁用默认事务操作
			Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
		},
	)
	if err != nil {
		panic(err)
	}
}

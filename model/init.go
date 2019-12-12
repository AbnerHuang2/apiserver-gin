package model

import (
	"apidemo-gin/conf"
	"fmt"
	"github.com/jinzhu/gorm"
)
// DataBase 用来组织数据库信息，实际使用中可能会有Master和Slave主从库
type DataBase struct {
	Master *gorm.DB
}

var DB *DataBase

func (dataBase *DataBase) Init() {
	cfg := conf.Cfg.Database
	_ = &DataBase{
		Master: openDB(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname),
	}
}

func openDB(user, password, host, port, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 设置日志格式和连接池大小
	db.LogMode(conf.Cfg.Database.LogMode)
	db.DB().SetMaxOpenConns(conf.Cfg.Database.MaximumPoolSize)
	db.DB().SetMaxIdleConns(conf.Cfg.Database.MaximumIdleSize)

	return db
}

func (dataBase *DataBase) Close()  {
	_ = DB.Master.Close()
}

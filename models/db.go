package models

import (
	"fmt"

	"ad.cherry.cn/pkg/logging"
	"ad.cherry.cn/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB database
var DB *gorm.DB

func init() {
	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		logging.Warn(err)
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 全局设置表名不可以为复数形式
	// db.SingularTable(true)

	DB = db
}

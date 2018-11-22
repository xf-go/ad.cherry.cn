package setting

import (
	"log"

	"github.com/go-ini/ini"
)

// App app结构体
type App struct {
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// AppSetting App设置
var AppSetting = &App{}

// Database 数据库连接结构体
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting Database设置
var DatabaseSetting = &Database{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("app", AppSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}
}

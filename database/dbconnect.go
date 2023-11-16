package database

import (
	"algoliteos/logger"
	"algoliteos/mvc"
	"database/sql"
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // init only
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mattn/go-sqlite3"

	"algoliteos/config"
)

const sqlite3Go = "sqlite3_with_go_func"
const admin = "admin"
const recordNotFound = "record not found"

func init() {
	sql.Register(sqlite3Go, &sqlite3.SQLiteDriver{
		ConnectHook: func(c *sqlite3.SQLiteConn) error {
			return c.RegisterFunc("regexp", regexp.MatchString, true)
		},
	})
}

var DB *gorm.DB

type DBUtil struct {
	db *gorm.DB
}

func GetDBUtil(db *gorm.DB) *DBUtil {
	return &DBUtil{db: db}
}

func InitDB() {
	var err error

	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	dbPath := v.GetString("db.path")
	conf.Unlock()

	sqlDb, err := sql.Open(sqlite3Go, dbPath)
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open("sqlite3", sqlDb)
	if err != nil {
		logger.Error("open %s failed, error: %s", dbPath, err)
		panic(err)
	} else {
		DB.DB().SetMaxOpenConns(1)
	}

	DB.SingularTable(true)

	_ = GetDBUtil(DB).CreateTableIfNotExist(&mvc.Record{}, "record_id", "id")
	_ = GetDBUtil(DB).CreateTableIfNotExist(&mvc.AlgoTaskSql{}, "task_id", "id")
}

// 创建表, 支持索引
func (d *DBUtil) CreateTableIfNotExist(schema interface{}, indexName string, columns ...string) error {
	var db = d.db
	if !db.HasTable(schema) {
		if err := db.Debug().CreateTable(schema).Error; err != nil {
			return err
		}
		// 添加唯一索引
		if indexName != "" || len(columns) != 0 {
			db.Model(schema).AddIndex("idx_"+indexName, columns...)
		}
	} else {
		db.AutoMigrate(schema)
	}
	return nil
}

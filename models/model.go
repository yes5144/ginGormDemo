package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

// Db xxx
var Db *gorm.DB

// BaseModel xxx
type BaseModel struct {
	// gorm.Model
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt string `gorm:"column:create_time" json:"create_time"`
	UpdatedAt string `gorm:"column:update_time" json:"update_time"`
}

// InitDb xxx
func InitDb() *gorm.DB {
	var err error
	// some var
	dbType := viper.GetString("datasource.driverName")
	dbuser := viper.GetString("datasource.dbuser")
	dbpass := viper.GetString("datasource.dbpass")
	dbhost := viper.GetString("datasource.dbhost")
	dbport := viper.GetString("datasource.dbport")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	tablePrefix := viper.GetString("datasource.tablePrefix")

	// connInfo := "root:channel@tcp(192.168.204.222:3306)/ginGormDemo?charset=utf8&parseTime=True&loc=Local"
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbuser, dbpass, dbhost, dbport, dbname, charset)

	log.Println(dbType, connInfo)
	// db, err = gorm.Open(dbType, connInfo)
	if dbType != "mysql" {
		connInfo = "sqlite3.db"
	}
	Db, err = gorm.Open(dbType, connInfo)
	if err != nil {
		log.Printf("gorm.Open db err, Failed code %#v", err)
		panic(fmt.Sprintf("gorm.Open db err, Failed code %v", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.AutoMigrate(&User{})
	return Db
}

// NowTime xxx
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// BeforeCreate xxx
func (v BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", NowTime())
	scope.SetColumn("update_time", NowTime())
	return nil
}

// BeforeUpdate xxx
func (v BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", NowTime())
	return nil
}

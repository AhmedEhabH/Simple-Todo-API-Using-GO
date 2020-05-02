package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB is a variable to connect with database
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig configuration with database
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:   "0.0.0.0",
		Port:   3306,
		User:   "root",
		DBName: "todos",
		//Password: "", // I think for me not need to password
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

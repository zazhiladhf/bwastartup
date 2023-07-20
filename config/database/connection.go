package database

import (
	// "AntarJemput-Be-C/models"
	"bwastartup/config"
	"fmt"

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(appConfig *config.Config) *gorm.DB {
	maxIdleConn := config.GetInt("DB_MAX_IDLE_CONNECTIONS", 10)
	maxConn := config.GetInt("DB_MAX_CONNECTIONS", 100)
	maxLifetimeConn := config.GetInt("DB_MAX_LIFETIME_CONNECTIONS", 2)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		appConfig.DBUsername,
		appConfig.DBPassword,
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBName,
	)

	// dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// InitCreateTable(db)
	log.Println("database connect success 🚀🚀🚀")
	log.Println("dsn :", connStr)
	return db

}

// func InitCreateTable(db *gorm.DB) {
// 	db.AutoMigrate(
// 		&models.Province{},
// 		&models.Regency{},
// 		&models.District{},
// 		&models.Agents{},
// 		&models.Customer{},
// 		&models.Users{},
// 		&models.ServiceTransaction{},
// 		&models.Transactions{},
// 		&models.TypeServiceTransaction{},
// 	)

// }

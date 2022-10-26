package driver

import (
	"dbservice/model"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	_ = godotenv.Load()
}

func GetDBConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	//set idle and max conn
	postgres, err := db.DB()
	if err != nil {
		return nil, err
	}
	idleConn, err := strconv.ParseInt(os.Getenv("DB_IDLE_CONN"), 10, 64)
	if err != nil {
		return nil, err
	}
	maxConn, err := strconv.ParseInt(os.Getenv("DB_IDLE_CONN"), 10, 64)
	if err != nil {
		return nil, err
	}
	if idleConn != 0 {
		postgres.SetMaxIdleConns(100)
	}
	if maxConn != 0 {
		postgres.SetMaxOpenConns(100)
	}

	//automigrate
	err = AutoMigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Event{},
		model.Register{},
		model.Subscription{},
		model.Like{},
	)
}

func InsertFK(db *gorm.DB) error {
	return nil
}

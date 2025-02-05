package configuration

import (
	"farhan/entity"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	username := os.Getenv("DATASOURCE_USERNAME")
	password := os.Getenv("DATASOURCE_PASSWORD")
	host := os.Getenv("DATASOURCE_HOST")
	port := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATASOURCE_DB_NAME")

	maxPoolOpen, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_MAX_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	maxPoolIdle, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_IDLE_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_LIFE_TIME"))
	if err != nil {
		log.Fatal(err)
	}

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxPollLifeTime) * time.Millisecond)

	// AutoMigrate
	err = db.AutoMigrate(
		&entity.Transaction{},
		&entity.Account{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

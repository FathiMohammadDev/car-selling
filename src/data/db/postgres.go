package db

import (
	"fmt"
	"log"

	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=admin dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	var err error
	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()

	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)

	logger.Info(logging.Postgres, logging.Startup, "DB conected successfuly.", nil)
	log.Println("DB conected successfuly.")

	return nil

}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	println("close DB")
	con, _ := dbClient.DB()
	con.Close()
}

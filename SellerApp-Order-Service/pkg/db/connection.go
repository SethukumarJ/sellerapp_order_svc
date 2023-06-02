package db

import (
	"fmt"

	"github.com/SethukumarJ/sellerapp-order-svc/pkg/config"
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	fmt.Println("database", db)

	db.AutoMigrate(
		domain.Order{},
		domain.Item{},
	)


	return db, dbErr
}

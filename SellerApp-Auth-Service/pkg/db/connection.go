package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/SethukumarJ/sellerapp-auth-service/pkg/config"
	domain "github.com/SethukumarJ/sellerapp-auth-service/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", "postgres", "postgres", "test", "5432", "1111")
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{})

	return db, dbErr
}

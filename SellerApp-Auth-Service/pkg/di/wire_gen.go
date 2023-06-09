// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/api"
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/api/handler"
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/config"
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/db"
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/repository"
	"github.com/SethukumarJ/sellerapp-auth-service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	jwtUsecase := usecase.NewJWTUsecase()
	userHandler := handler.NewUserHandler(userUseCase, jwtUsecase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}

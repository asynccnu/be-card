// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/asynccnu/be-card/grpc"
	"github.com/asynccnu/be-card/ioc"
	"github.com/asynccnu/be-card/pkg/grpcx"
	"github.com/asynccnu/be-card/repository"
	"github.com/asynccnu/be-card/repository/cache"
	"github.com/asynccnu/be-card/repository/dao"
	"github.com/asynccnu/be-card/service"
)

// Injectors from wire.go:

func InitGRPCServer() grpcx.Server {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	daoDao := dao.NewCardDao(db)
	cmdable := ioc.InitRedis()
	cacheCache := cache.NewCardRedisCache(cmdable)
	cardRepository := repository.NewCardRepository(daoDao, cacheCache)
	serviceService := service.NewCardService(cardRepository)
	cardService := grpc.NewCardGrpcService(serviceService)
	client := ioc.InitEtcdClient()
	server := ioc.InitGRPCxKratosServer(cardService, client, logger)
	return server
}

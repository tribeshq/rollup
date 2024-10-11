// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/tribeshq/tribes/configs"
	"github.com/tribeshq/tribes/internal/domain/entity"
	"github.com/tribeshq/tribes/internal/infra/cartesi/handler/advance_handler"
	"github.com/tribeshq/tribes/internal/infra/cartesi/handler/inspect_handler"
	"github.com/tribeshq/tribes/internal/infra/cartesi/middleware"
	"github.com/tribeshq/tribes/internal/infra/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewMiddlewares() (*Middlewares, error) {
	gormDB, err := configs.SetupSQlite()
	if err != nil {
		return nil, err
	}
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	rbacMiddleware := middleware.NewRBACMiddleware(userRepositorySqlite)
	middlewares := &Middlewares{
		RBAC: rbacMiddleware,
	}
	return middlewares, nil
}

func NewMiddlewaresMemory() (*Middlewares, error) {
	gormDB, err := configs.SetupSQliteMemory()
	if err != nil {
		return nil, err
	}
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	rbacMiddleware := middleware.NewRBACMiddleware(userRepositorySqlite)
	middlewares := &Middlewares{
		RBAC: rbacMiddleware,
	}
	return middlewares, nil
}

func NewAdvanceHandlers() (*AdvanceHandlers, error) {
	gormDB, err := configs.SetupSQlite()
	if err != nil {
		return nil, err
	}
	bidRepositorySqlite := db.NewBidRepositorySqlite(gormDB)
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	contractRepositorySqlite := db.NewContractRepositorySqlite(gormDB)
	auctionRepositorySqlite := db.NewAuctionRepositorySqlite(gormDB)
	bidAdvanceHandlers := advance_handler.NewBidAdvanceHandlers(bidRepositorySqlite, userRepositorySqlite, contractRepositorySqlite, auctionRepositorySqlite)
	userAdvanceHandlers := advance_handler.NewUserAdvanceHandlers(userRepositorySqlite, contractRepositorySqlite)
	auctionAdvanceHandlers := advance_handler.NewAuctionAdvanceHandlers(bidRepositorySqlite, userRepositorySqlite, auctionRepositorySqlite, contractRepositorySqlite)
	contractAdvanceHandlers := advance_handler.NewContractAdvanceHandlers(contractRepositorySqlite)
	advanceHandlers := &AdvanceHandlers{
		BidAdvanceHandlers:      bidAdvanceHandlers,
		UserAdvanceHandlers:     userAdvanceHandlers,
		AuctionAdvanceHandlers:  auctionAdvanceHandlers,
		ContractAdvanceHandlers: contractAdvanceHandlers,
	}
	return advanceHandlers, nil
}

func NewAdvanceHandlersMemory() (*AdvanceHandlers, error) {
	gormDB, err := configs.SetupSQliteMemory()
	if err != nil {
		return nil, err
	}
	bidRepositorySqlite := db.NewBidRepositorySqlite(gormDB)
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	contractRepositorySqlite := db.NewContractRepositorySqlite(gormDB)
	auctionRepositorySqlite := db.NewAuctionRepositorySqlite(gormDB)
	bidAdvanceHandlers := advance_handler.NewBidAdvanceHandlers(bidRepositorySqlite, userRepositorySqlite, contractRepositorySqlite, auctionRepositorySqlite)
	userAdvanceHandlers := advance_handler.NewUserAdvanceHandlers(userRepositorySqlite, contractRepositorySqlite)
	auctionAdvanceHandlers := advance_handler.NewAuctionAdvanceHandlers(bidRepositorySqlite, userRepositorySqlite, auctionRepositorySqlite, contractRepositorySqlite)
	contractAdvanceHandlers := advance_handler.NewContractAdvanceHandlers(contractRepositorySqlite)
	advanceHandlers := &AdvanceHandlers{
		BidAdvanceHandlers:      bidAdvanceHandlers,
		UserAdvanceHandlers:     userAdvanceHandlers,
		AuctionAdvanceHandlers:  auctionAdvanceHandlers,
		ContractAdvanceHandlers: contractAdvanceHandlers,
	}
	return advanceHandlers, nil
}

func NewInspectHandlers() (*InspectHandlers, error) {
	gormDB, err := configs.SetupSQlite()
	if err != nil {
		return nil, err
	}
	bidRepositorySqlite := db.NewBidRepositorySqlite(gormDB)
	bidInspectHandlers := inspect_handler.NewBidInspectHandlers(bidRepositorySqlite)
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	contractRepositorySqlite := db.NewContractRepositorySqlite(gormDB)
	userInspectHandlers := inspect_handler.NewUserInspectHandlers(userRepositorySqlite, contractRepositorySqlite)
	auctionRepositorySqlite := db.NewAuctionRepositorySqlite(gormDB)
	auctionInspectHandlers := inspect_handler.NewAuctionInspectHandlers(auctionRepositorySqlite)
	contractInspectHandlers := inspect_handler.NewContractInspectHandlers(contractRepositorySqlite)
	inspectHandlers := &InspectHandlers{
		BidInspectHandlers:      bidInspectHandlers,
		UserInspectHandlers:     userInspectHandlers,
		AuctionInspectHandlers:  auctionInspectHandlers,
		ContractInspectHandlers: contractInspectHandlers,
	}
	return inspectHandlers, nil
}

func NewInspectHandlersMemory() (*InspectHandlers, error) {
	gormDB, err := configs.SetupSQliteMemory()
	if err != nil {
		return nil, err
	}
	bidRepositorySqlite := db.NewBidRepositorySqlite(gormDB)
	bidInspectHandlers := inspect_handler.NewBidInspectHandlers(bidRepositorySqlite)
	userRepositorySqlite := db.NewUserRepositorySqlite(gormDB)
	contractRepositorySqlite := db.NewContractRepositorySqlite(gormDB)
	userInspectHandlers := inspect_handler.NewUserInspectHandlers(userRepositorySqlite, contractRepositorySqlite)
	auctionRepositorySqlite := db.NewAuctionRepositorySqlite(gormDB)
	auctionInspectHandlers := inspect_handler.NewAuctionInspectHandlers(auctionRepositorySqlite)
	contractInspectHandlers := inspect_handler.NewContractInspectHandlers(contractRepositorySqlite)
	inspectHandlers := &InspectHandlers{
		BidInspectHandlers:      bidInspectHandlers,
		UserInspectHandlers:     userInspectHandlers,
		AuctionInspectHandlers:  auctionInspectHandlers,
		ContractInspectHandlers: contractInspectHandlers,
	}
	return inspectHandlers, nil
}

// wire.go:

var setBidRepositoryDependency = wire.NewSet(db.NewBidRepositorySqlite, wire.Bind(new(entity.BidRepository), new(*db.BidRepositorySqlite)))

var setAuctionRepositoryDependency = wire.NewSet(db.NewAuctionRepositorySqlite, wire.Bind(new(entity.AuctionRepository), new(*db.AuctionRepositorySqlite)))

var setContractRepositoryDependency = wire.NewSet(db.NewContractRepositorySqlite, wire.Bind(new(entity.ContractRepository), new(*db.ContractRepositorySqlite)))

var setUserRepositoryDependency = wire.NewSet(db.NewUserRepositorySqlite, wire.Bind(new(entity.UserRepository), new(*db.UserRepositorySqlite)))

var setAdvanceHandlers = wire.NewSet(advance_handler.NewBidAdvanceHandlers, advance_handler.NewUserAdvanceHandlers, advance_handler.NewAuctionAdvanceHandlers, advance_handler.NewContractAdvanceHandlers)

var setInspectHandlers = wire.NewSet(inspect_handler.NewBidInspectHandlers, inspect_handler.NewUserInspectHandlers, inspect_handler.NewAuctionInspectHandlers, inspect_handler.NewContractInspectHandlers)

var setMiddleware = wire.NewSet(middleware.NewRBACMiddleware)

type Middlewares struct {
	RBAC *middleware.RBACMiddleware
}

type AdvanceHandlers struct {
	BidAdvanceHandlers      *advance_handler.BidAdvanceHandlers
	UserAdvanceHandlers     *advance_handler.UserAdvanceHandlers
	AuctionAdvanceHandlers  *advance_handler.AuctionAdvanceHandlers
	ContractAdvanceHandlers *advance_handler.ContractAdvanceHandlers
}

type InspectHandlers struct {
	BidInspectHandlers      *inspect_handler.BidInspectHandlers
	UserInspectHandlers     *inspect_handler.UserInspectHandlers
	AuctionInspectHandlers  *inspect_handler.AuctionInspectHandlers
	ContractInspectHandlers *inspect_handler.ContractInspectHandlers
}
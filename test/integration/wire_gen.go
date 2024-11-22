// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tribeshq/tribes/internal/domain/entity"
	"github.com/tribeshq/tribes/internal/infra/cartesi/handler/advance_handler"
	"github.com/tribeshq/tribes/internal/infra/cartesi/handler/inspect_handler"
	"github.com/tribeshq/tribes/internal/infra/cartesi/middleware"
	"github.com/tribeshq/tribes/internal/infra/repository"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func NewMiddlewares(gormDB *gorm.DB) (*Middlewares, error) {
	userRepositorySqlite := repository.NewUserRepositorySqlite(gormDB)
	tlsnMiddleware := middleware.NewTLSNMiddleware(userRepositorySqlite)
	rbacMiddleware := middleware.NewRBACMiddleware(userRepositorySqlite)
	middlewares := &Middlewares{
		TLSN: tlsnMiddleware,
		RBAC: rbacMiddleware,
	}
	return middlewares, nil
}

func NewAdvanceHandlers(gormDB *gorm.DB) (*AdvanceHandlers, error) {
	userRepositorySqlite := repository.NewUserRepositorySqlite(gormDB)
	orderRepositorySqlite := repository.NewOrderRepositorySqlite(gormDB)
	contractRepositorySqlite := repository.NewContractRepositorySqlite(gormDB)
	crowdfundingRepositorySqlite := repository.NewCrowdfundingRepositorySqlite(gormDB)
	orderAdvanceHandlers := advance_handler.NewOrderAdvanceHandlers(userRepositorySqlite, orderRepositorySqlite, contractRepositorySqlite, crowdfundingRepositorySqlite)
	userAdvanceHandlers := advance_handler.NewUserAdvanceHandlers(userRepositorySqlite, contractRepositorySqlite)
	crowdfundingAdvanceHandlers := advance_handler.NewCrowdfundingAdvanceHandlers(orderRepositorySqlite, userRepositorySqlite, crowdfundingRepositorySqlite, contractRepositorySqlite)
	contractAdvanceHandlers := advance_handler.NewContractAdvanceHandlers(contractRepositorySqlite)
	advanceHandlers := &AdvanceHandlers{
		OrderAdvanceHandlers:        orderAdvanceHandlers,
		UserAdvanceHandlers:         userAdvanceHandlers,
		CrowdfundingAdvanceHandlers: crowdfundingAdvanceHandlers,
		ContractAdvanceHandlers:     contractAdvanceHandlers,
	}
	return advanceHandlers, nil
}

func NewInspectHandlers(gormDB *gorm.DB) (*InspectHandlers, error) {
	orderRepositorySqlite := repository.NewOrderRepositorySqlite(gormDB)
	orderInspectHandlers := inspect_handler.NewOrderInspectHandlers(orderRepositorySqlite)
	userRepositorySqlite := repository.NewUserRepositorySqlite(gormDB)
	contractRepositorySqlite := repository.NewContractRepositorySqlite(gormDB)
	userInspectHandlers := inspect_handler.NewUserInspectHandlers(userRepositorySqlite, contractRepositorySqlite)
	crowdfundingRepositorySqlite := repository.NewCrowdfundingRepositorySqlite(gormDB)
	crowdfundingInspectHandlers := inspect_handler.NewCrowdfundingInspectHandlers(crowdfundingRepositorySqlite)
	contractInspectHandlers := inspect_handler.NewContractInspectHandlers(contractRepositorySqlite)
	inspectHandlers := &InspectHandlers{
		OrderInspectHandlers:        orderInspectHandlers,
		UserInspectHandlers:         userInspectHandlers,
		CrowdfundingInspectHandlers: crowdfundingInspectHandlers,
		ContractInspectHandlers:     contractInspectHandlers,
	}
	return inspectHandlers, nil
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(repository.NewOrderRepositorySqlite, wire.Bind(new(entity.OrderRepository), new(*repository.OrderRepositorySqlite)))

var setCrowdfundingRepositoryDependency = wire.NewSet(repository.NewCrowdfundingRepositorySqlite, wire.Bind(new(entity.CrowdfundingRepository), new(*repository.CrowdfundingRepositorySqlite)))

var setContractRepositoryDependency = wire.NewSet(repository.NewContractRepositorySqlite, wire.Bind(new(entity.ContractRepository), new(*repository.ContractRepositorySqlite)))

var setUserRepositoryDependency = wire.NewSet(repository.NewUserRepositorySqlite, wire.Bind(new(entity.UserRepository), new(*repository.UserRepositorySqlite)))

var setAdvanceHandlers = wire.NewSet(advance_handler.NewOrderAdvanceHandlers, advance_handler.NewUserAdvanceHandlers, advance_handler.NewCrowdfundingAdvanceHandlers, advance_handler.NewContractAdvanceHandlers)

var setInspectHandlers = wire.NewSet(inspect_handler.NewOrderInspectHandlers, inspect_handler.NewUserInspectHandlers, inspect_handler.NewCrowdfundingInspectHandlers, inspect_handler.NewContractInspectHandlers)

var setMiddleware = wire.NewSet(middleware.NewTLSNMiddleware, middleware.NewRBACMiddleware)

type Middlewares struct {
	TLSN *middleware.TLSNMiddleware
	RBAC *middleware.RBACMiddleware
}

type AdvanceHandlers struct {
	OrderAdvanceHandlers        *advance_handler.OrderAdvanceHandlers
	UserAdvanceHandlers         *advance_handler.UserAdvanceHandlers
	CrowdfundingAdvanceHandlers *advance_handler.CrowdfundingAdvanceHandlers
	ContractAdvanceHandlers     *advance_handler.ContractAdvanceHandlers
}

type InspectHandlers struct {
	OrderInspectHandlers        *inspect_handler.OrderInspectHandlers
	UserInspectHandlers         *inspect_handler.UserInspectHandlers
	CrowdfundingInspectHandlers *inspect_handler.CrowdfundingInspectHandlers
	ContractInspectHandlers     *inspect_handler.ContractInspectHandlers
}
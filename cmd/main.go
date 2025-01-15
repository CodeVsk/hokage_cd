//go:debug x509negativeserial=1
package main

import (
	httpserver "github.com/codevsk/codevsk_golang_cd/internal/infra/httpServer"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/httpserver/handler"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/migration"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/repository"
)

func main() {
	db, err := orm.NewGormSqlServerDatabase()
	if err != nil {
		panic(err)
	}

	if err := migration.AutoMigrate(db); err != nil {
		panic(err)
	}

	applicationRepository := repository.NewApplicationRepository(db)

	server := httpserver.NewHttpServer(":9090")

	applicationHandler := handler.NewApplicationHandler(applicationRepository)

	server.RegisterHandler("POST", "create", applicationHandler.Create)

	server.Start()
}
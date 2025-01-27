//go:debug x509negativeserial=1
package main

import (
	http "github.com/codevsk/codevsk_golang_cd/internal/infra/httpServer"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/httpserver/handler"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/migration"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
)

func main() {
	db, err := orm.NewGormSqlServerDatabase()
	if err != nil {
		panic(err)
	}

	if err := migration.AutoMigrate(db); err != nil {
		panic(err)
	}


	unitOfWork, err := uow.NewUnitOfWork(db)
	if err != nil {
		panic(err)
	}

	server := http.NewHttpServer(":9090")

	applicationHandler := handler.NewApplicationHandler(unitOfWork)
	
	server.RegisterHandler("application/:id", "GET", applicationHandler.GetById)
	server.RegisterHandler("application/create", "POST", applicationHandler.Create)
	server.RegisterHandler("application/:id", "PATCH", applicationHandler.Update)
	server.RegisterHandler("application/:id", "DELETE", applicationHandler.Delete)
	
	deploymentHandler := handler.NewDeploymentHandler(unitOfWork)

	server.RegisterHandler("deployment/publish", "POST", deploymentHandler.Publish)

	server.Start()
}

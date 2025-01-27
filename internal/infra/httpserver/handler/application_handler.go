package handler

import (
	"net/http"

	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	usecase "github.com/codevsk/codevsk_golang_cd/internal/usecase/application"
	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct{
	Uow uow.UnitOfWork
}

func NewApplicationHandler(uow uow.UnitOfWork) *ApplicationHandler {
	return &ApplicationHandler{
		Uow: uow,
	}
}

func (a *ApplicationHandler) Create(c *gin.Context) {
	var input dto.CreateApplicationInputDTO
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	u := usecase.NewCreateApplicationUseCase(a.Uow)
	result := u.Execute(input);

	c.JSON(int(result.StatusCode), result.Body)
}

func (a *ApplicationHandler) Update(c *gin.Context) {
	var input dto.UpdateApplicationInputDTO
	applicationId := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	u := usecase.NewUpdateApplicationUseCase(a.Uow)
	result := u.Execute(applicationId, input);

	c.JSON(int(result.StatusCode), result.Body)
}

func (a *ApplicationHandler) Delete(c *gin.Context) {
	applicationId := c.Param("id")

	u := usecase.NewDeleteApplicationUseCase(a.Uow)
	result := u.Execute(applicationId);

	c.JSON(int(result.StatusCode), result.Body)
}

func (a *ApplicationHandler) GetById(c *gin.Context) {
	applicationId := c.Param("id")

	u := usecase.NewGetByIdApplicationUseCase(a.Uow)
	result := u.Execute(applicationId);

	c.JSON(int(result.StatusCode), result.Body)
}
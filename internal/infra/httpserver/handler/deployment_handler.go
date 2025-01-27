package handler

import (
	"net/http"

	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	usecase "github.com/codevsk/codevsk_golang_cd/internal/usecase/deployment"
	"github.com/gin-gonic/gin"
)

type DeploymentHandler struct{
	Uow uow.UnitOfWork
}

func NewDeploymentHandler(uow uow.UnitOfWork) *DeploymentHandler {
	return &DeploymentHandler{
		Uow: uow,
	}
}

func (a *DeploymentHandler) Publish(c *gin.Context) {
	ctx := c.Request.Context()

	var input dto.PublishApplicationInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	u := usecase.NewPublishDeploymentUseCase(ctx, a.Uow)
	result := u.Execute(input);

	c.JSON(int(result.StatusCode), result.Body)
}
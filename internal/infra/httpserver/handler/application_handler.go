package handler

import (
	"net/http"

	"github.com/codevsk/codevsk_golang_cd/internal/contract"
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	usecase "github.com/codevsk/codevsk_golang_cd/internal/usecase/application/create"
	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct{
	ApplicationRepository contract.ApplicationRepository
}

func NewApplicationHandler(ApplicationRepository contract.ApplicationRepository) *ApplicationHandler {
	return &ApplicationHandler{
		ApplicationRepository: ApplicationRepository,
	}
}

func (a *ApplicationHandler) Create(c *gin.Context) {
	var input dto.CreateApplicationInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	createApplication := usecase.NewCreateApplicationUseCase(a.ApplicationRepository)
	result := createApplication.Execute(input);

	c.JSON(int(result.StatusCode), result.Body)
}
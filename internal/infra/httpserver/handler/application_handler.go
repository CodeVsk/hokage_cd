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
		c.String(http.StatusBadRequest, "Body is invalid!")
		return
	}
	
	createApplication := usecase.NewCreateApplicationUseCase(a.ApplicationRepository)
	if err := createApplication.Execute(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application created succeffully!"})
}
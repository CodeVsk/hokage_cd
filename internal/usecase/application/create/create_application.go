package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/contract"
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
)

type CreateApplicationUseCase struct {
	ApplicationRepository contract.ApplicationRepository
}

func NewCreateApplicationUseCase(ApplicationRepository contract.ApplicationRepository) *CreateApplicationUseCase {
	return &CreateApplicationUseCase{
		ApplicationRepository: ApplicationRepository,
	}
}

func (c *CreateApplicationUseCase) Execute(input dto.CreateApplicationInputDTO) (result.Result) {
	application , err := entity.NewApplication(input.Name, input.Slug, input.Repository_url, "master")
	if(err != nil){
		return result.NewValidationResult(err.Error())
	}

	if _, err := c.ApplicationRepository.GetBySlug(application.Slug); err == nil {
		return result.NewConflictResult()
	}

	model := model.NewApplication(application)

	if err := c.ApplicationRepository.Create(model); err != nil {
		return result.NewInternalResult(err)
	}

	return result.NewCreatedResult()
}
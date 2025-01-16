package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/contract"
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
)

type CreateApplicationUseCase struct {
	ApplicationRepository contract.ApplicationRepository
}

func NewCreateApplicationUseCase(ApplicationRepository contract.ApplicationRepository) *CreateApplicationUseCase {
	return &CreateApplicationUseCase{
		ApplicationRepository: ApplicationRepository,
	}
}

func (c *CreateApplicationUseCase) Execute(input dto.CreateApplicationInputDTO) error {
	application , err := entity.NewApplication(input.Name, input.Slug, input.Repository_url, "master")
	if(err != nil){
		return err
	}

	model := model.NewApplication(application)

	if err := c.ApplicationRepository.Create(model); err != nil {
		return err
	}

	return nil
}
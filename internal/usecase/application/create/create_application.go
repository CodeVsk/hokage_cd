package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/contract"
	"github.com/codevsk/codevsk_golang_cd/internal/dto"
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
	return nil
}
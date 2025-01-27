package usecase

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
)

type GetByIdApplicationUseCase struct {
	Uow uow.UnitOfWork
}

func NewGetByIdApplicationUseCase(uow uow.UnitOfWork) *GetByIdApplicationUseCase {
	return &GetByIdApplicationUseCase{
		Uow: uow,
	}
}

func (c *GetByIdApplicationUseCase) Execute(id string) (result.Result) {
	application, err := c.Uow.GetApplicationRepository().GetById(id)
	if err != nil {
		return result.NewInternalResult(err)
	}

	return result.NewSuccessResult(application, 0, 0)
}
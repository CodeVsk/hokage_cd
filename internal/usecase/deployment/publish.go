package deployment

import (
	"context"

	"github.com/codevsk/codevsk_golang_cd/internal/dto"
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/uow"
	"github.com/codevsk/codevsk_golang_cd/pkg/result"
)

type PublishDeploymentUseCase struct {
	Uow uow.UnitOfWork
	Ctx context.Context
}

func NewPublishDeploymentUseCase(ctx context.Context, uow uow.UnitOfWork) *PublishDeploymentUseCase {
	return &PublishDeploymentUseCase{
		Uow: uow,
		Ctx: ctx,
	}
}

func (c *PublishDeploymentUseCase) Execute(input dto.PublishApplicationInputDTO) result.Result {

	//git.GetFileContentFromGitHub(c.Ctx, inpu)

	//application, err := entity.NewApplication(input.Name, input.Slug, input.Repository_url, "master")
	//if err != nil {
	//	return result.NewValidationResult(err.Error())
	//}
//
	//if _, err := c.Uow.GetApplicationRepository().GetBySlug(application.Slug); err == nil {
	//	return result.NewConflictResult()
	//}
//
	//model := model.ToApplicationModel(application)
//
	//if err := c.Uow.GetApplicationRepository().Create(model); err != nil {
	//	c.Uow.Rollback()
//
	//	return result.NewInternalResult(err)
	//}
//
	//if err := c.Uow.Commit(); err != nil {
	//	return result.NewInternalResult(err)
	//}

	return result.NewCreatedResult()
}
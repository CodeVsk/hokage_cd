package contract

import (
	"github.com/codevsk/codevsk_golang_cd/internal/infra/orm/model"
)

type ApplicationRepository interface {
	Create(model *model.Application) error
}

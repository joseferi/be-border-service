package usecase

import (
	"be-border-service/internal/common"
)

type UseCase interface {
	Serve(data *common.Data) common.Response
}

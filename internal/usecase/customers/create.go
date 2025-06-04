package customers

import (
	"be-border-service/internal/common"
	"be-border-service/internal/repository"
	"be-border-service/internal/usecase"
	"net/http"
)

type customerUseCase struct {
	findOneCustomer repository.FindOneCustomer
}

func NewCustomerUseCase(findOneCustomer repository.FindOneCustomer) usecase.UseCase {
	return &customerUseCase{
		findOneCustomer: findOneCustomer,
	}
}

func (c *customerUseCase) Serve(data *common.Data) common.Response {
	return *common.NewResponse().WithStatusCode(http.StatusOK).WithMessage("test")
}

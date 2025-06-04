package customers

import (
	"be-border-service/internal/common"
	"be-border-service/internal/delivery/customers/request"
	"be-border-service/internal/repository"
	"be-border-service/internal/usecase"
	"be-border-service/pkg/logger"
	"fmt"
	"net/http"
)

type retrieveCustomerUCase struct {
	repo repository.FindAllCustomer
}

func NewRetrieveCustomerUseCase(repo repository.FindAllCustomer) usecase.UseCase {
	return &retrieveCustomerUCase{
		repo: repo,
	}
}

func (rc *retrieveCustomerUCase) Serve(r *common.Data) common.Response {
	var (
		ctx    = r.Request.Context()
		params request.CustomerQueryParams
		lf     = logger.NewFields(
			logger.EventName("ucase.retrieve.customers"),
			logger.Any("http.request.method", r.Request.Method),
			logger.Any("http.request.url", r.Request.URL),
		)
	)
	if err := r.Cast(&params); err != nil {
		return *common.NewResponse().WithStatusCode(http.StatusBadRequest).WithMessage("invalid request")
	}
	lf.Append(logger.Any("params", params))
	customers, err := rc.repo.FindAll(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error find all customers, got err :%v", err), lf...)
		return *common.NewResponse().
			WithStatusCode(http.StatusInternalServerError).
			WithMessage("something went wrong")
	}
	logger.Info("Success retrieve data customers", lf...)
	return *common.NewResponse().WithStatusCode(http.StatusOK).
		WithMessage("success retrieve data customers").WithData(customers)
}

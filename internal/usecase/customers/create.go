package customers

import (
	"be-border-service/internal/common"
	"be-border-service/internal/delivery/customers/request"
	"be-border-service/internal/model"
	"be-border-service/internal/repository"
	"be-border-service/internal/usecase"
	"be-border-service/internal/validators"
	"be-border-service/pkg/logger"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type customerUseCase struct {
	find  repository.FindOneCustomer
	store repository.CreateCustomer
}

func NewCreateCustomerUseCase(findOneCustomer repository.FindOneCustomer, store repository.CreateCustomer) usecase.UseCase {
	return &customerUseCase{
		find:  findOneCustomer,
		store: store,
	}
}

func (c *customerUseCase) Serve(r *common.Data) common.Response {
	var (
		payload request.CreateCustomerRequest
		ctx     = r.Request.Context()
		lf      = logger.NewFields(
			logger.EventName("ucase.retrieve.customers"),
			logger.Any("http.request.method", r.Request.Method),
			logger.Any("http.request.url", r.Request.URL),
		)
	)
	if err := r.Cast(&payload); err != nil {
		logger.Warn(fmt.Sprintf("error cast payload got err :%v", err), lf...)
		return *common.NewResponse().WithStatusCode(http.StatusBadRequest).WithMessage("invalid payload")
	}
	if err := validator.New().Struct(payload); err != nil {
		msg := validators.FormatValidationErrors(err, payload)
		logger.Warn(fmt.Sprintf("error validate payload got %v", msg), lf...)
		return *common.NewResponse().
			WithStatusCode(http.StatusUnprocessableEntity).
			WithMessage("validation error").
			WithError(msg)
	}
	var err error
	customer, err := c.find.FindOne(ctx, model.Users{
		Email: payload.Email,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed find customer got err :%v", err), lf...)
		return *common.NewResponse().WithStatusCode(http.StatusInternalServerError).WithMessage("something went wrong")
	}
	if customer != nil {
		logger.Warn("customer already exist", lf...)
		return *common.NewResponse().WithStatusCode(http.StatusConflict).WithMessage("customer already exist")
	}
	lf.Append(logger.Any("payload", payload))

	if err = c.store.Create(ctx, model.Users{
		Name:  payload.Name,
		Email: payload.Email,
	}); err != nil {
		logger.Error(fmt.Sprintf("error store customer got err :%v", err), lf...)
		return *common.NewResponse().WithStatusCode(http.StatusInternalServerError).WithMessage("something went wrong")
	}
	logger.Info("success create customer", lf...)
	return *common.NewResponse().WithStatusCode(http.StatusOK).WithMessage("ok")
}

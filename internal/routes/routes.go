package routes

import (
	"be-border-service/internal/bootstrap"
	"be-border-service/internal/common"
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"be-border-service/internal/handler"
	"be-border-service/internal/middleware"
	"be-border-service/internal/repository"
	"be-border-service/internal/usecase"
	"be-border-service/internal/usecase/customers"
	"be-border-service/pkg/logger"
	"be-border-service/pkg/workerx"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

type router struct {
	*mux.Router
	config *config.Config
}

func NewRouter(cfg *config.Config) Router {
	return &router{
		Router: mux.NewRouter(),
		config: cfg,
	}
}
func (rtr *router) handle(hfn httpHandlerFunc, svc usecase.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
				w.WriteHeader(http.StatusInternalServerError)
				res := common.Response{
					Status: http.StatusInternalServerError,
				}

				logger.Error(fmt.Sprintf("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res.Byte())

				return
			}
		}()

		// validate middleware
		if !middleware.FilterFunc(w, r, rtr.config, mdws) {
			return
		}

		resp := hfn(r, svc, rtr.config)
		rtr.response(w, resp)
	}
}

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp common.Response) {
	w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
	w.WriteHeader(resp.Status)
	w.Write(resp.Byte())
	return
}

func (r *router) Route() *router {
	r.Router.NotFoundHandler = http.HandlerFunc(middleware.NotFound)
	db := bootstrap.RegistryDatabase(&r.config.DB)

	root := r.Router.PathPrefix("/").Subrouter()
	v1 := root.PathPrefix("/v1").Subrouter()

	// init workersClient
	workerClient := workerx.NewAsynqClient(r.config)

	// initialize repo
	customerRepo := repository.NewCustomerRepository(db)

	// initialize usecase
	createCustomer := customers.NewCreateCustomerUseCase(customerRepo, customerRepo, workerClient)
	retrieveCustomer := customers.NewRetrieveCustomerUseCase(customerRepo)

	healthCheck := usecase.NewHealthCheck(workerClient)

	r.HandleFunc("/healthcheck", r.handle(
		handler.HttpRequest,
		healthCheck,
		middleware.HealthCheckMiddleware(),
	)).Methods(http.MethodGet)

	v1.HandleFunc("/customer", r.handle(
		handler.HttpRequest,
		createCustomer,
	)).Methods(http.MethodPost)

	v1.HandleFunc("/customer", r.handle(
		handler.HttpRequest,
		retrieveCustomer,
	)).Methods(http.MethodGet)
	return r
}

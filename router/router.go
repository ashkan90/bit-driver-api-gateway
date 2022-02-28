package router

import (
	"github.com/ashkan90/bit-driver-api-gateway/config"
	"github.com/ashkan90/bit-driver-api-gateway/handler"
	middleware2 "github.com/ashkan90/bit-driver-api-gateway/middleware"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

type ProxyRouter struct {
	Logger *log.Logger
	Config *config.GeneralConfig
}

func (pr *ProxyRouter) NewRouter() *mux.Router {
	pr.Logger.Println("Proxy router has been registered...")
	var sv = mux.NewRouter()
	var methods = []string{
		http.MethodHead,
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
	}

	for _, service := range pr.Config.Services {
		pr.Logger.Println("= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")
		pr.Logger.Println("Service name is ", service.Name)
		pr.Logger.Println("Service target is ", service.Target)
		pr.Logger.Println("Service path is ", service.Path)
		pr.Logger.Println("= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = \n")

		var proxy = &httputil.ReverseProxy{
			Director: handler.NewHandler(service.Target, service.Path),
		}
		var selectedStrategy = (middleware2.Strategy)(service.Strategy)
		var strategyExecutor = pr.strategySelector(proxy, selectedStrategy)

		sv.PathPrefix(service.Listen).HandlerFunc(strategyExecutor).Methods(methods...)

		//
		sv.PathPrefix(service.Listen).HandlerFunc(middleware2.FwdOptions(proxy)).Methods(http.MethodOptions)
	}

	return sv
}

func (*ProxyRouter) strategySelector(proxy *httputil.ReverseProxy, strategy middleware2.Strategy) http.HandlerFunc {
	switch strategy {
	case middleware2.StrategyCheckAuth:
		return middleware2.CheckAuth(proxy)
	default:
		return middleware2.FwdOptions(proxy)
	}
}

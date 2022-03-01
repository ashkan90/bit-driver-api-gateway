package router

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/ashkan90/bit-driver-api-gateway/config"
	"github.com/ashkan90/bit-driver-api-gateway/handler"
	middleware_strategy "github.com/ashkan90/bit-driver-api-gateway/middleware"

	"github.com/gorilla/mux"
)

type ProxyRouter struct {
	Logger *log.Logger
	Config *config.GeneralConfig
}

// NewRouter declares all given services as proxy connection.
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
		pr.Logger.Println("= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")
		pr.Logger.Println()

		var proxy = &httputil.ReverseProxy{
			Director: handler.NewHandler(service.Target, service.Path),
		}
		var selectedStrategy = (middleware_strategy.Strategy)(service.Strategy)
		var strategyExecutor = pr.strategySelector(proxy, selectedStrategy)

		sv.PathPrefix(service.Listen).HandlerFunc(strategyExecutor).Methods(methods...).Name(service.Path + " " + service.Strategy)

		// Some browsers can send a request with OPTION method
		// and in other case It'll be blocked cause of CheckAuth strategy
		// to bypass that, we implemented a Forward strategy tho.
		sv.PathPrefix(service.Listen).HandlerFunc(middleware_strategy.FwdOptions(proxy)).Methods(http.MethodOptions).Name(service.Path + " fwd")
	}

	return sv
}

// strategySelector determines which strategy will be executed with given parameter
func (*ProxyRouter) strategySelector(proxy *httputil.ReverseProxy, strategy middleware_strategy.Strategy) http.HandlerFunc {
	switch strategy {
	case middleware_strategy.StrategyCheckAuth:
		return middleware_strategy.CheckAuth(proxy)
	case middleware_strategy.StrategyForwardDirectly:
		return middleware_strategy.FwdOptions(proxy)
	default:
		return middleware_strategy.FwdOptions(proxy)
	}
}

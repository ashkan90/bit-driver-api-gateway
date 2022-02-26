package router

import (
	"github.com/ashkan90/bit-driver-api-gateway/proxy/config"
	"github.com/ashkan90/bit-driver-api-gateway/proxy/handler"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/ashkan90/bit-driver-api-gateway/proxy/middleware"
	"github.com/gorilla/mux"
)

type ProxyRouter struct {
	Logger  *log.Logger
	Config  *config.ServiceConfig
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

	sv.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("asdasd"))
		writer.WriteHeader(200)
	})

	for _, service := range pr.Config.Services {
		pr.Logger.Println("Service target is ", service.Target)
		pr.Logger.Println("Service path is ", service.Path)

		var proxy = &httputil.ReverseProxy{
			Director:       handler.NewHandler(service.Target),
		}
		var selectedStrategy = (middleware.Strategy)(service.Strategy)
		var strategyExecutor = pr.strategySelector(proxy, selectedStrategy)

		sv.PathPrefix(service.Path).HandlerFunc(strategyExecutor).Methods(methods...)

		//
		sv.PathPrefix(service.Path).HandlerFunc(middleware.FwdOptions(proxy)).Methods(http.MethodOptions)
	}

	return sv
}

func (*ProxyRouter) strategySelector(proxy *httputil.ReverseProxy, strategy middleware.Strategy) http.HandlerFunc {
	switch strategy {
	case middleware.StrategyCheckAuth:
		return middleware.CheckAuth(proxy)
	default:
		return middleware.FwdOptions(proxy)
	}
}
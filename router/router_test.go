package router

import (
	"github.com/ashkan90/bit-driver-api-gateway/config"
	"github.com/ashkan90/bit-driver-api-gateway/middleware"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProxyRouter_NewRouterIsInitiatingMuxCorrectly(t *testing.T) {
	var service = config.Service{
		Name:     "test-svc-1",
		Target:   "localhost:1010",
		Strategy: string(middleware.StrategyCheckAuth),
		Listen:   "/listen-path/",
		Path:     "/target-path",
	}
	var router = &ProxyRouter{
		Logger: log.Default(),
		Config: &config.GeneralConfig{
			Services: []config.Service{
				service,
			},
		},
	}

	m := router.NewRouter()

	var routeNames []string
	_ = m.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeNames = append(routeNames, route.GetName())
		return nil
	})

	assert.Contains(t, routeNames, service.Path+" "+service.Strategy)
	assert.Contains(t, routeNames, service.Path+" fwd")
}

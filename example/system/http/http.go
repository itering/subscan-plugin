package http

import (
	"github.com/itering/subscan-plugin/example/system/service"
	"github.com/itering/subscan-plugin/router"
	"net/http"
)

var (
	svc *service.Service
)

func Router(s *service.Service) []router.Http {
	svc = s
	return []router.Http{
		{"test", system},
	}
}

func system(w http.ResponseWriter, r *http.Request) error {
	return nil
}

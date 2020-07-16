package router

import "net/http"

type Http struct {
	Router string
	Handle func(w http.ResponseWriter, r *http.Request) error
}

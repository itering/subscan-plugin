package router

import "net/http"

// Http router
type Http struct {
	// Router path
	Router string

	// Router response handle
	Handle func(w http.ResponseWriter, r *http.Request) error
}

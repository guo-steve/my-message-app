package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
)

func InitRounter(handler *Handler) http.Handler {
	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(handler.logger))
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(handler.Authenticator(
			handler.services.AuthService,
		))

		r.Route("/v1/messages", func(r chi.Router) {
			r.Post("/", handler.PostMessage) // POST /messages
			r.Get("/", handler.GetMessages)  // GET /messages
		})

		r.Route("/v1/logout", func(r chi.Router) {
			r.Post("/", handler.Logout) // POST /logout
		})
	})

	// public routes
	router.Group(func(r chi.Router) {
		r.Get("/v1/ping", handler.Ping)
		r.Post("/v1/login", handler.Login)
		r.Post("/v1/register", handler.Register)
	})

	return router
}

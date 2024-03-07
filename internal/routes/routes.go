package routes

import (
	"flc_backend/internal/notifications"
	"flc_backend/internal/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	router.Get("/api/v1/healthz", healthzHandler)

	//users
	router.Post("/api/v1/createUser", users.CreateUserHandler)
	router.Post("/api/v1/userLogin", users.UserLoginHandler)

	//get notices
	router.Get("/api/v1/notifications", notifications.GetNotificationsHandler)

	return router
}

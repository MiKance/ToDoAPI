package handlers

import (
	"github.com/MiKance/ToDoAPI/internal/handlers/auth"
	"github.com/MiKance/ToDoAPI/internal/handlers/items"
	"github.com/MiKance/ToDoAPI/internal/handlers/lists"
	"github.com/MiKance/ToDoAPI/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *chi.Mux {
	authHandl := auth.NewAuthHandler(h.service)
	listHandl := lists.NewToDoListHandler(h.service)
	itemHandl := items.NewItemHandler(h.service)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", authHandl.SignIn())
		r.Post("/sign-up", authHandl.SignUp())
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(h.userIdentify)
		r.Route("/lists", func(r chi.Router) {
			r.Post("/", listHandl.CreateList())
			r.Get("/", listHandl.GetLists())
			r.Get("/{id}", listHandl.GetListByID())
			r.Put("/{id}", listHandl.UpdateList())
			r.Delete("/{id}", listHandl.DeleteList())

			r.Route("/{list_id}/items", func(r chi.Router) {
				r.Post("/", itemHandl.CreateItem())
				r.Get("/", itemHandl.GetItemByListId())
				r.Get("/{id}", itemHandl.GetItemByID())
				r.Put("/{id}", itemHandl.UpdateItem())
				r.Delete("/{id}", itemHandl.Delete())
			})

		})

	})

	return router
}

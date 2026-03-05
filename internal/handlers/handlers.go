package handlers

import (
	"github.com/MiKance/ToDoAPI/internal/handlers/auth"
	"github.com/MiKance/ToDoAPI/internal/handlers/items"
	"github.com/MiKance/ToDoAPI/internal/handlers/lists"
	"github.com/MiKance/ToDoAPI/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

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

			r.Route("/{id}/items", func(r chi.Router) {
				r.Post("/", items.CreateItem())
				r.Get("/", items.GetItems())
				r.Get("/{id}", items.GetItemByID())
				r.Put("/{id}", items.UpdateItem())
				r.Delete("/{id}", items.Delete())
			})

		})

	})

	return router
}

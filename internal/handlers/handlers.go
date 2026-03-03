package handlers

import (
	"ToDoAPI/internal/handlers/auth"
	"ToDoAPI/internal/handlers/items"
	"ToDoAPI/internal/handlers/lists"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
}

func (h *Handler) InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", auth.SignIn())
		r.Post("/sign-up", auth.SignUp())
	})

	router.Route("/api", func(r chi.Router) {
		r.Route("/lists", func(r chi.Router) {
			r.Post("/", lists.CreateList())
			r.Get("/", lists.GetLists())
			r.Get("/{id}", lists.GetListByID())
			r.Put("/{id}", lists.UpdateList())
			r.Delete("/{id}", lists.DeleteList())

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

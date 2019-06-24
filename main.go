package main

import (
	"Expenses_REST-API/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)


func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/expenses", func(r chi.Router) {
		r.Get("/", handlers.ListExpenses)
		r.Post("/", handlers.CreateExpense)

		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", handlers.ListOneExpense)
			r.Put("/", handlers.UpdateExpense)
			r.Delete("/", handlers.DeleteExpense)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}




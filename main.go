package main

import (
	"Expenses_REST-API/Interfaces"
	"Expenses_REST-API/expenseDB"
	"Expenses_REST-API/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func handleRequests(db Interfaces.Database){
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/expenses", func(r chi.Router) {
		r.Get("/", db.GetAll)
		r.Post("/", db.CreateExpense)

		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", db.GetId)
			r.Put("/", db.UpdateExpense)
			r.Delete("/", db.DeleteExpense)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	mongodb :=  &handlers.MongoDB{Db: expenseDB.ExpCollFunc()}
	handleRequests(mongodb)
}




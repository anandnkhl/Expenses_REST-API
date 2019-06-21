package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"google.golang.org/genproto/googleapis/type/date"
	"log"
	"net/http"
)

type Expense struct{
	Id 				int 		`json:"id"`
	Description 	string 		`json:"description"`
	Type 			string 		`json:"type"`
	Amount 			float64 	`json:"amount"`
	CreatedOn 		date.Date 	`json:"created_on"`
	UpdatedOn 		date.Date 	`json:"updated_on"`
}

type ExpenseRequest struct{
	*Expense
}

type ExpenseResponse struct{
	*Expense
}

var expense Expense
//type Expenses []Expense

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/expenses", func(r chi.Router) {
		//r.Get("/", ListExpenses)
		r.Post("/", CreateExpense)

		//r.Route("/{articleID}", func(r chi.Router) {
		//	r.Use(ArticleCtx)
		//	r.Get("/", ListOneExpense)
		//	r.Put("/", UpdateExpense)
		//	r.Delete("/", DeleteExpense)
		//})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}



func CreateExpense(writer http.ResponseWriter, request *http.Request) {
	data := &ExpenseRequest{}
	err := render.Bind(request, data)
	if err != nil {
		_ = render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	expense = *data.Expense

	_ = render.Render(writer, request, NewExpenseResponse(&expenses))

}



func NewExpenseResponse(expense *Expense) *ExpenseResponse {
	resp := &ExpenseResponse{Expense: expense}
	return resp
}


func (Expense) Bind(request *http.Request) error {
	return nil
}

type ErrResponse struct {
	Err error `json:"-"`
	HTTPStatusCode int `json:"-"`
	StatusText string `json:"status"`
}

func (ExpenseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err: err,
		HTTPStatusCode: 400,
		StatusText: "Invalid Request",
	}
}
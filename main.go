package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Expense struct{
	Id 				int 		`json:"id"`
	Description 	string 		`json:"description"`
	Type 			string 		`json:"type"`
	Amount 			float64 	`json:"amount"`
	CreatedOn 		time.Time 	`json:"created_on"`
	UpdatedOn 		time.Time 	`json:"updated_on"`
}

type Expenses []Expense

type ExpenseRequest struct{
	*Expense
}

type ExpenseResponse struct{
	*Expense
}

type ExpensesResponse struct {
	*Expenses
}


var expense Expense
var expenses Expenses

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/expenses", func(r chi.Router) {
		r.Get("/", ListExpenses)
		r.Post("/", CreateExpense)

		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", ListOneExpense)
			r.Put("/", UpdateExpense)
			r.Delete("/", DeleteExpense)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func UpdateExpense(writer http.ResponseWriter, request *http.Request) {
	var data ExpenseRequest
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	for index,exp := range expenses{
		if exp.Id == ID{
			idTemp := expenses[index].Id
			createdOnTemp := expenses[index].CreatedOn

			_ = render.Bind(request, &data)

			expenses[index] = *data.Expense
			expenses[index].UpdatedOn = time.Now()
			expenses[index].CreatedOn = createdOnTemp
			expenses[index].Id = idTemp
			return
		}
	}
}

func DeleteExpense(writer http.ResponseWriter, request *http.Request) {

	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	for index,exp := range expenses{
		if exp.Id == ID{
			expenses = append(expenses[:index], expenses[index + 1 :]... )
			return
		}
	}
}


func ListOneExpense(writer http.ResponseWriter, request *http.Request) {
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	for _,exp := range expenses{
		if exp.Id == ID{
			_ = render.Render(writer, request, NewExpenseResponse(&exp))
			return
		}
	}
}

func ListExpenses(writer http.ResponseWriter, request *http.Request) {
	_=render.Render(writer, request, AllExpensesResponse(&expenses))
}


func (ExpensesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func AllExpensesResponse(exp *Expenses) *ExpensesResponse {
	resp := &ExpensesResponse{Expenses: exp}
	return resp
}


func CreateExpense(writer http.ResponseWriter, request *http.Request) {
	data := &ExpenseRequest{}
	err := render.Bind(request, data)
	if err != nil {
		_ = render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	expense = *data.Expense
	expense.Id = len(expenses) + 1
	expense.CreatedOn = time.Now()
	expense.UpdatedOn = time.Now()
	expenses = append(expenses, expense)
	_ = render.Render(writer, request, NewExpenseResponse(&expense))

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
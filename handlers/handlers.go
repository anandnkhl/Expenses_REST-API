package handlers

import (
	"Expenses_REST-API/types"
	"github.com/go-chi/render"
	"net/http"
)

type ExpensesResponse struct {
	*types.Expenses
}

func (ExpensesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func AllExpensesResponse(exp *types.Expenses) *ExpensesResponse {
	resp := &ExpensesResponse{Expenses: exp}
	return resp
}


type ExpenseResponse struct{
	*types.Expense
}

func NewExpenseResponse(expense *types.Expense) *ExpenseResponse {
	resp := &ExpenseResponse{Expense: expense}
	return resp
}

type CreateExpenseRequest struct{
	*types.Expense
}

func (CreateExpenseRequest) Bind(request *http.Request) error {
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
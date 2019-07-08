package handlers

import (
	"Expenses_REST-API/types"
	"net/http"
)

//go:generate ../autogen -dbtype=MongoDB -op=Create

type CreateExpenseRequest struct{
	*types.Expense
}

func NewExpenseResponse(expense *types.Expense) *ExpenseResponse {
	resp := &ExpenseResponse{Expense: expense}
	return resp
}

type ExpenseResponse struct{
	*types.Expense
}

func (ExpenseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (CreateExpenseRequest) Bind(request *http.Request) error {
	return nil
}
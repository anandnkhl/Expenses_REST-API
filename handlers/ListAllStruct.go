package handlers

import (
	"Expenses_REST-API/types"
	"net/http"
)

//go:generate ../autogen -dbtype=MongoDB -op=ListAll

type ExpensesResponse struct {
	*types.Expenses
}

func AllExpensesResponse(exp *types.Expenses) *ExpensesResponse {
	resp := &ExpensesResponse{Expenses: exp}
	return resp
}

func (ExpensesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
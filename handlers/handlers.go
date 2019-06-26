package handlers

import (
	"Expenses_REST-API/types"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)
type MongoDB struct{
	Db *mongo.Collection
}




func (mongo *MongoDB)GetId(writer http.ResponseWriter, request *http.Request) {
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	curr:= mongo.Db.FindOne(ctx, bson.D{{"id", ID}})
	var expense types.Expense
	_ = curr.Decode(&expense)
	_=render.Render(writer, request, NewExpenseResponse(&expense))
}



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
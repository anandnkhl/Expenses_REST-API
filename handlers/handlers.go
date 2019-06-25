package handlers

import (
	"Expenses_REST-API/expenseDB"
	"Expenses_REST-API/types"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)
var expense types.Expense
var expenses types.Expenses

func CreateExpense(writer http.ResponseWriter, request *http.Request) {
	data := &CreateExpenseRequest{}
	err := render.Bind(request, data)
	if err != nil {
		_ = render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	expense = *data.Expense
	expense.Id = len(expenses) + 1
	expense.CreatedOn = time.Now().String()
	expense.UpdatedOn = time.Now().String()
	expenses = append(expenses, expense)

	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	_, _ = expenseDB.ExpCollFunc().InsertOne(ctx, expense)
	_ = render.Render(writer, request, NewExpenseResponse(&expense))
}

func DeleteExpense(writer http.ResponseWriter, request *http.Request) {

	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))

	filter := bson.D{{"id", ID}}
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	expenseDB.ExpCollFunc().FindOneAndDelete(ctx, filter)
}


func UpdateExpense(writer http.ResponseWriter, request *http.Request) {
	var data CreateExpenseRequest
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))

	for index,exp := range expenses{
		if exp.Id == ID{
			idTemp := expenses[index].Id
			createdOnTemp := expenses[index].CreatedOn

			_ = render.Bind(request, &data)

			expenses[index] = *data.Expense
			expenses[index].UpdatedOn = time.Now().String()
			expenses[index].CreatedOn = createdOnTemp
			expenses[index].Id = idTemp
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
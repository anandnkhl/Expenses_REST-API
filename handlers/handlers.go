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

func (mongo *MongoDB)CreateExpense(writer http.ResponseWriter, request *http.Request) {

	timeIDString := time.Now().String()
	timeIDString = timeIDString[:4]+timeIDString[5:7]+timeIDString[8:10]+timeIDString[11:13]+
		timeIDString[14:16]+timeIDString[17:19]+timeIDString[21:23]
	timeIDInt,_ := strconv.Atoi(timeIDString)

	data := &CreateExpenseRequest{}
	err := render.Bind(request, data)
	if err != nil {
		_ = render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	expense := *data.Expense
	expense.Id = timeIDInt
	expense.CreatedOn = time.Now().String()
	expense.UpdatedOn = time.Now().String()

	ctx,_ := context.WithTimeout(context.Background(), 15*time.Second)
	_, _ = mongo.Db.InsertOne(ctx, expense)

	_ = render.Render(writer, request, NewExpenseResponse(&expense))
}

func (mongo *MongoDB)DeleteExpense(writer http.ResponseWriter, request *http.Request) {

	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))

	filter := bson.D{{"id", ID}}
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	mongo.Db.FindOneAndDelete(ctx, filter)
}


func (mongo *MongoDB)UpdateExpense(writer http.ResponseWriter, request *http.Request) {
	var data CreateExpenseRequest
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	ctx := context.TODO()
	curr:= mongo.Db.FindOne(ctx, bson.D{{"id", ID}})
	_ = curr.Decode(&data)

	_ = render.Bind(request, &data)

	update := bson.D{
		{"$set", bson.D{{"id",ID}}},
		{"$set", bson.D{{"description", data.Description}}},
		{"$set", bson.D{{"type", data.Type}}},
		{"$set", bson.D{{"amount", data.Amount}}},
		{"$set", bson.D{{"updated_on", time.Now().String()}}},
	}

	_, _ = mongo.Db.UpdateOne(ctx, bson.D{{"id", ID}}, update,)

}


func (mongo *MongoDB)GetId(writer http.ResponseWriter, request *http.Request) {
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	curr:= mongo.Db.FindOne(ctx, bson.D{{"id", ID}})
	var expense types.Expense
	_ = curr.Decode(&expense)
	_=render.Render(writer, request, NewExpenseResponse(&expense))
}

func (mongo *MongoDB)GetAll(writer http.ResponseWriter, request *http.Request) {
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	curr,_ := mongo.Db.Find(ctx, bson.D{})
	var expenses types.Expenses
	for curr.Next(ctx){
		var expense types.Expense
		_ = curr.Decode(&expense)
		expenses = append(expenses, expense)
	}
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
package handlers

import (
	"Expenses_REST-API/types"
	"context"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func (mongo *MongoDB)GetAll(writer http.ResponseWriter, request *http.Request) {
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	curr,_ := mongo.Db.Find(ctx, bson.D{})
	var expenses types.Expenses
	for curr.Next(ctx){
		var expense types.Expense
		_ = curr.Decode(&expense)
		expenses = append(expenses, expense)
	}
	_ = render.Render(writer, request, AllExpensesResponse(&expenses))
}
package handlers

import (
	"Expenses_REST-API/types"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)

func (mongo *MongoDB)GetId(writer http.ResponseWriter, request *http.Request) {
	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	curr:= mongo.Db.FindOne(ctx, bson.D{{"id", ID}})
	var expense types.Expense
	_ = curr.Decode(&expense)
	_=render.Render(writer, request, NewExpenseResponse(&expense))
}

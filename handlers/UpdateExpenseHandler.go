package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)

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

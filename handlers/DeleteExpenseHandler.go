package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)

func (mongo *MongoDB)DeleteExpense(writer http.ResponseWriter, request *http.Request) {

	ID, _ := strconv.Atoi(chi.URLParam(request, "ID"))

	filter := bson.D{{"id", ID}}
	ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
	mongo.Db.FindOneAndDelete(ctx, filter)
}
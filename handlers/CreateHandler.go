package handlers

import (
    "context"
    "github.com/go-chi/render"
    "net/http"
    "strconv"
    "time"
)

func ( m *MongoDB ) CreateExpense (writer http.ResponseWriter, request *http.Request){

    timeIDStr := time.Now().String()
    timeIDStr = timeIDStr[:4]+timeIDStr[5:7]+timeIDStr[8:10]+timeIDStr[11:13]+ timeIDStr[14:16]+timeIDStr[17:19]+timeIDStr[21:23]
    timeIDInt,_ := strconv.Atoi(timeIDStr)

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

    ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)
    _, _ = m.Db.InsertOne(ctx, expense)
    //TODO implement the renderer
    _ = render.Render(writer, request, NewExpenseResponse(&expense))
}
package handlers

import (
    "context"
    "net/http"
    "time"
)

func ( m *MongoDB ) UpdateExpense (writer http.ResponseWriter, request *http.Request){
    ctx,_ := context.WithTimeout(context.Background(), 5*time.Second)

    //TODO implement the renderer
    //_ = render.Render(writer, request, RendererFunc(&obj))
}
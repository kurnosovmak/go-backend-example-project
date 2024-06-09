package routes

import (
	"encoding/json"
	"example/cmd/routes/requests"
	"example/modules/example/infra/adapter"
	"fmt"
	"net/http"
)

const RequestParseError = "Request bad format"
const ServerErrorText = "Server error Please retry"

var exampleAdapter *adapter.ExampleAdapter

func RegisterRoutes() {
	http.HandleFunc("/example.example", apiExampleExample)
}

func init() {
	exampleAdapter = adapter.Create()
}

func apiExampleExample(response http.ResponseWriter, r *http.Request) {
	// Сначало получаем request
	request := requests.ApiExampleExampleRequest{}
	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		// тут нужно логировать
		http.Error(response, RequestParseError, http.StatusBadRequest)
		return
	}

	// Запрашиваем нашу логику
	text, err := exampleAdapter.PushHelloWorld(request.Name)
	if err != nil {
		// тут нужно логировать
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	// Заполняем результат
	_, err = fmt.Fprintln(response, text)

	if err != nil {
		// тут нужно логировать
		http.Error(response, ServerErrorText, http.StatusInternalServerError)
		return
	}
}

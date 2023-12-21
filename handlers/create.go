package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/todogo/models"
)

func Create(response http.ResponseWriter, request *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro decode do json %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao tentar inserir todo ao banco: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inseriro com sucesso! ID: %d", id),
		}
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)
}

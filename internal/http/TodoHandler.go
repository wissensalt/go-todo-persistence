package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/wissensalt/go-todo-persistence/internal/repository"
	"github.com/wissensalt/go-todo-persistence/internal/service"
	"net/http"
	"strconv"
)

type (
	TodoHandler interface {
		Create(writer http.ResponseWriter, request *http.Request)
		Update(writer http.ResponseWriter, request *http.Request)
		Get(writer http.ResponseWriter, request *http.Request)
		GetById(writer http.ResponseWriter, request *http.Request)
		Delete(writer http.ResponseWriter, request *http.Request)
	}

	TodoHandlerImpl struct {
		service.TodoServiceImpl
	}
)

func (t TodoHandlerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	setContentTypeToJson(writer)
	var todo repository.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
	}

	err = t.TodoServiceImpl.Create(todo)
	if err != nil {
		http.Error(writer, "Failed to create todo", http.StatusInternalServerError)
	}
}

func (t TodoHandlerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	setContentTypeToJson(writer)
	var todo repository.Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
	}

	err = t.TodoServiceImpl.Update(todo)
	if err != nil {
		http.Error(writer, "Failed to update todo", http.StatusInternalServerError)
	}
}

func (t TodoHandlerImpl) Get(writer http.ResponseWriter, request *http.Request) {
	setContentTypeToJson(writer)
	todos := t.TodoServiceImpl.Get()
	err := json.NewEncoder(writer).Encode(todos)
	if err != nil {
		http.Error(writer, "Failed write response", http.StatusInternalServerError)
	}
}

func (t TodoHandlerImpl) GetById(writer http.ResponseWriter, request *http.Request) {
	setContentTypeToJson(writer)
	idParam := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(idParam)
	todo := t.TodoServiceImpl.GetById(id)
	err := json.NewEncoder(writer).Encode(&todo)
	if err != nil {
		http.Error(writer, "Failed write response", http.StatusInternalServerError)
	}
}

func (t TodoHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	setContentTypeToJson(writer)
	idParam := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(idParam)
	err := t.TodoServiceImpl.Delete(id)
	if err != nil {
		http.Error(writer, "Failed to delete todo", http.StatusInternalServerError)
	}
}

func setContentTypeToJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

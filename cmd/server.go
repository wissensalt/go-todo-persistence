package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/wissensalt/go-todo-persistence/internal/config"
	api "github.com/wissensalt/go-todo-persistence/internal/http"
	"github.com/wissensalt/go-todo-persistence/internal/repository"
	"github.com/wissensalt/go-todo-persistence/internal/service"
	"net/http"
)

var sqlDB *sql.DB

func main() {
	sqlDB = config.ConnectDB()
	r := chi.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Golang Todo Application with Persistence DB"))
	})
	r.Mount("/todos", TodoRouter())
	_ = http.ListenAndServe("localhost:8080", r)
}

func TodoRouter() http.Handler {
	todoRepository := repository.TodoRepositoryImpl{DB: sqlDB}
	todoService := service.TodoServiceImpl{TodoRepository: todoRepository}
	todoHandler := api.TodoHandlerImpl{TodoServiceImpl: todoService}
	r := chi.NewRouter()
	r.Get("/", todoHandler.Get)
	r.Get("/{id}", todoHandler.GetById)
	r.Post("/", todoHandler.Create)
	r.Put("/", todoHandler.Update)
	r.Delete("/{id}", todoHandler.Delete)

	return r
}

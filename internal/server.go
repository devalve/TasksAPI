package internal

import (
	"TasksAPI/internal/task"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(addr string) *http.Server {
	taskService := task.NewService()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", task.CreateTaskHandler(taskService)).Methods("POST")
	router.HandleFunc("/tasks/{id}", task.GetTaskHandler(taskService)).Methods("GET")
	router.HandleFunc("/tasks/{id}", task.DeleteTaskHandler(taskService)).Methods("DELETE")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

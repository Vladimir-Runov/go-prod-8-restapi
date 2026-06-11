package handlers

import (
	"encoding/json"
	"go-prod-8-restapi/internal/models"
	"go-prod-8-restapi/internal/storage"
	"net/http"
	"strconv"
)

type Handler struct{ Store storage.Storage }

func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

// /tasks (GET, POST)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	// TODO: реализуйте разбор метода, JSON, коды статусов, валидацию
	switch r.Method {

	case http.MethodGet:
		tasks := h.Store.List()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdTask, err := h.Store.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTask)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	// TODO: извлечение id, маршрутизация по методу, ошибки
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedTask, err := h.Store.Update(id, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedTask)

	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

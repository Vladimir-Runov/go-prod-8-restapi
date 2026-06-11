package handlers

import (
	"encoding/json"
	"fmt"
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
		tasks := h.Store.List() // Получаем список всех задач из хранилища
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTask)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	// TODO: извлечение id, маршрутизация по методу, ошибки
	idStr := r.URL.Path[len("/tasks/"):] // Извлекаем ID задачи из URL
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("invalid task id: %s", idStr), "message": "task id must be a valid integer"}) //
		return
	}

	switch r.Method {

	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			//http.Error(w, "task not found", http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("invalid task id: %s", idStr), "message": "task id not found"}) //
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("invalid PUT-body: %s", r.Body), "message": "invalid request body"}) //
			return
		}

		updatedTask, err := h.Store.Update(id, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("PUT-body: %s", r.Body), "message": "error update task"}) //
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedTask)

	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			//http.Error(w, err.Error(), http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("body: %s", r.Body), "message": "error update task"}) //

			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		//http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Method: %s", r.Method), "message": "method not allowed"}) //

	}
}

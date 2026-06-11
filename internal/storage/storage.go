package storage

import (
	"fmt"
	"go-prod-8-restapi/internal/models"
	"sync"
)

type Storage interface {
	List() []models.Task
	Create(models.Task) (models.Task, error)
	Get(id int) (models.Task, bool)
	Update(id int, tsk models.Task) (models.Task, error)
	Delete(id int) error
}

type Memory struct {
	tasks  map[int]models.Task
	mu     sync.RWMutex
	nextID int
}

func New() *Memory {
	return &Memory{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (m *Memory) Create(task models.Task) (models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = m.nextID
	m.tasks[m.nextID] = task
	m.nextID++

	return task, nil
}

func (m *Memory) List() []models.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	taskList := make([]models.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func (m *Memory) Get(id int) (models.Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, exists := m.tasks[id]
	return task, exists
}

func (m *Memory) Update(id int, task models.Task) (models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tasks[id]; !exists {
		return models.Task{}, fmt.Errorf("task with id %d not found", id)
	}

	task.ID = id // Сохраняем ID
	m.tasks[id] = task
	return task, nil
}

func (m *Memory) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tasks[id]; !exists {
		return fmt.Errorf("task with id %d not found", id)
	}

	delete(m.tasks, id)
	return nil
}


func Init(store Storage) { {
	
	createdTask1, err := store.Create(models.Task{Title: "задача - I", Done: false})
	if err != nil {
		fmt.Println("Ошибка при создании задачи 1:", err)
		return
	}
	fmt.Println("Создана задача 1:", createdTask1)
	createdTask2, err := store.Create(models.Task{Title: "задача - II", Done: false})
	if err != nil {
		fmt.Println("Ошибка при создании задачи 2:", err)
		return
	}
	fmt.Println("Создана задача 2:", createdTask2)
}
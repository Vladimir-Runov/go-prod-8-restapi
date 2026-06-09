package storage

import (
    "sync"
    "go-prod-8-restapi/internal/models"
)

type Storage interface {
    List() []models.Task
    Create(models.Task) (models.Task, error)
    Get(id int) (models.Task, bool)
    Update(id int, models.Task) (models.Task, error)
    Delete(id int) error
}


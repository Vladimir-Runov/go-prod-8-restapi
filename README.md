# go-prod-8-restapi
Создание REST API-сервиса

# Installation

```  
 cd путь_кпапке
 go work init
 go mod tidy
 go mod init goprod-8-dz-7
 go run main.go

```

goprod-8-restapi/  
  cmd/server/main.go  
  internal/handlers/tasks.go          // TODO: реализовать логику
  internal/models/task.go             // модель данных
  internal/storage/storage.go         // интерфейс Storage
  internal/http/middleware.go         // (опционально)
  README.md  
  go.mod  

```

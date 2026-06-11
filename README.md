# go-prod-8-restapi
Создание REST API-сервиса

# Installation

```  
 cd путь_кпапке
 go work init
 go mod init go-prod-8-restapi
 go mod tidy

 go run main.go
 go run ./cmd/server/main.go
 go build ./...
```
cd C:\Users\Admin\Documents\go\git_netology\go-prod-8-restapi

goprod-8-restapi/  

   go-prod-8-restapi/
   ├── cmd/
   │   └── server/
   │       └── main.go
   ├── internal/
   │   ├── handlers/  tasks.go          // TODO: реализовать логику
   │   ├── models/    task.go             // модель данных
   │   └── storage/   storage.go         // интерфейс Storage
   └── go.mod
    
       README.md  
       go.mod  
       test.cmd


```
http://localhost:8088
curl http://localhost:8088


 http://localhost:8088/exit
 
 или в терминале нажать Ctrl + C. Это отправит сигнал прерывания (SIGINT) вашему приложению, и сервер завершит свою работу.
 taskkill /PID <PID> /F

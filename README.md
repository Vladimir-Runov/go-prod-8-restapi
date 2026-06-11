# go-prod-8-restapi
Создание REST API-сервиса

# Installation

```  
 on main folder
 go work init
 go mod init go-prod-8-restapi
 go mod tidy

```
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
   │   ├── handlers/  tasks.go  
   │   ├── models/    task.go  
   │   └── storage/   storage.go  
   └── go.mod  
    
       README.md   
       go.mod  
       test.cmd  


```
http://localhost:8088
curl http://localhost:8088

### Завершить работу сервера:  
 http://localhost:8088/exit  
 или отправить сигнал прерывания (SIGINT): в терминале нажать Ctrl + C.   
 или taskkill /PID <PID> /F
```

### references:
https://github.com/golang/go/wiki/Modules
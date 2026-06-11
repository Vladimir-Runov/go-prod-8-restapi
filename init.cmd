cd ./internal/handlers
go mod init go-prod-8-restapi/internal/handlers
go mod tidy

cd ../models
go mod init go-prod-8-restapi/internal/models
go mod tidy

cd ../storage
go mod init go-prod-8-restapi/internal/storage
go mod tidy

cd ..\..
go vet ./...   

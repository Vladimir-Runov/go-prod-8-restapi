package main

import (
	"fmt"
	"go-prod-8-restapi/internal/handlers"
	"go-prod-8-restapi/internal/storage"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Port       string
	ServerName string
}

func exitHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is shutting down...")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shutting down server...")) // Send response before exiting
	os.Exit(0)                                 // Note: os.Exit will terminate the program immediately, so any deferred functions will not run.
}

func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%v] Request from r/a:%s url: %s", time.Now().Format(time.RFC3339), r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[%v] Request completed (%v ms) url: %s", time.Now().Format(time.RFC3339), time.Since(start), r.URL.Path)
	})
}

// aboutHandler обрабатывает запросы к /about и возвращает простое приветствие. Время ответа варьируется от 20 до 265 микросекунд, что имитирует некоторую нагрузку на сервер.
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(20+rand.Intn(246)) * time.Microsecond)
	fmt.Fprintf(w, "o nas!")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	config := Config{
		Port:       ":8088",
		ServerName: "srv-runov-001",
	}
	store := storage.New()
	storage.Init(store)
	h := handlers.New(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/exit", exitHandler)        //
	mux.HandleFunc("/about", aboutHandler)      //
	mux.HandleFunc("/tasks", h.TasksCollection) // GET, POST
	mux.HandleFunc("/tasks/", h.TaskItem)       // GET, PUT, DELETE

	loggedMux := timingMiddleware(mux) // Оборачиваем mux в middleware для логирования и измерения времени выполнения
	server := &http.Server{Addr: config.Port, Handler: loggedMux}

	go func() {
		log.Printf("[%s] Server listening on %s", config.ServerName, config.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs // Wait for a termination signal
	log.Printf("[%s] Shutting down server...", config.ServerName)
	if err := server.Close(); err != nil {
		log.Fatal(err) // Handle error during server shutdown
	}
}

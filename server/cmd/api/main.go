package main

import (
	"fmt"
	"money-backend/internal/handler"
	"money-backend/pkg/database"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! Ini server Go pertama saya")
}

func main() {
	database.InitDB()
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/transactions", handler.CreateTransaction)

	port := ":8000"
	fmt.Println("Server berjalan di http://localhost" + port)

	http.ListenAndServe(port, nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kmwk10/subscription-service/internal/config"
	"github.com/kmwk10/subscription-service/internal/db"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	database := db.Connect(cfg)
	defer database.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	log.Printf("Server is running on port %s\n", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, nil); err != nil {
		log.Fatal(err)
	}
}

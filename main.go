package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")

	admin_models.User{}.Migrate()
	admin_models.Post{}.Migrate()
	admin_models.Measurement{}.Migrate()
	admin_models.Workout{}.Migrate()
	admin_models.Diet{}.Migrate()

	log.Println("server started on :8080")
	log.Println("localhost:" + port + "/")

	router := config.Routes()
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}

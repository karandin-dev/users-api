package main

import (
	"log"
	"userapi_gin/iternal/users"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=usersdb sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	storage := users.NewSQLStorage(db)
	service := users.NewService(storage)
	handler := users.NewHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Println("server started at :8080")
	r.Run(":8080")
}

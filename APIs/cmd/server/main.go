package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andremarquezz/Go_Expert/APIS/configs"
	"github.com/andremarquezz/Go_Expert/APIS/internal/entity"
	"github.com/andremarquezz/Go_Expert/APIS/internal/infra/database"
	"github.com/andremarquezz/Go_Expert/APIS/internal/infra/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	port := ":8080"
	http.HandleFunc("/products", productHandler.CreateProduct)

	fmt.Fprintf(os.Stderr, "Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao iniciar o servidor: %v\n", err)
	}
}

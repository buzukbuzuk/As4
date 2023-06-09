package main

import (
	"clean-architecture/pkg/store/postgres"
	"clean-architecture/services/contact/internal/delivery"
	"clean-architecture/services/contact/internal/repository"
	"clean-architecture/services/contact/internal/usecase"
	"fmt"
	"net/http"
)

func main() {

	url := &postgres.ConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "interesting",
		Password: "1234",
		DbName:   "interesting-clean",
	}

	db, err := postgres.OpenDB(url)
	if err != nil {
		fmt.Printf("postgres.OpenDB: %v", err)
	}

	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("application started")

	http.ListenAndServe("localhost:4000", delivery.Mux)
}

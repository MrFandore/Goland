package main

import (
	httpx "github.com/MrFandore/Practica_11/internal/http"
	"github.com/MrFandore/Practica_11/internal/http/handlers"
	"github.com/MrFandore/Practica_11/internal/repo"
	"log"
	"net/http"
)

func main() {
	repo := repo.NewNoteRepoMem()
	h := &handlers.Handler{Repo: repo}
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

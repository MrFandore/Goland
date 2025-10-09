package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hero/practica2/utils"
)

type pingResp struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func Run() {
	mux := http.NewServeMux()

	// Корневой маршрут — текст
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Hello, Go project structure!")
	})

	// JSON-маршрут /ping
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(pingResp{
			Status: "ok",
			Time:   time.Now().UTC().Format(time.RFC3339),
		})
	})

	utils.LogInfo("Server is starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		utils.LogError("server error: " + err.Error())
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice-docker-network-web/repository"
)

func main() {
	server := http.Server{
		Addr: ":8888",
	}

	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			userInfos, err := repository.GetTodos()
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err), http.StatusInternalServerError)
			}

			if err := json.NewEncoder(w).Encode(userInfos); err != nil {
				http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err), http.StatusInternalServerError)
			}
		default:
			return
		}
	})
	server.ListenAndServe()
}

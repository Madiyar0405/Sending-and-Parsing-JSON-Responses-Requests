package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var requestBody RequestBody

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				http.Error(w, "Invalid JSON message", http.StatusBadRequest)
				return
			}

			if requestBody.Message != "" {
				fmt.Println("Received message:", requestBody.Message)

				response := ResponseBody{
					Status:  "success",
					Message: "Данные успешно приняты",
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				encoder := json.NewEncoder(w)
				encoder.Encode(response)
			} else {
				http.Error(w, "Invalid or missing 'message' field", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}

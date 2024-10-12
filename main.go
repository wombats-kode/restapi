package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type cmdresult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Go Home Simple REST API Server")
}

func getDate(w http.ResponseWriter, _ *http.Request) {
	result := cmdresult{}

	out, err := exec.Command("date").Output()
	if err == nil {
		result.Success = true
		result.Message = "The date is " + string(out)
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/v1/getdate", getDate)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}

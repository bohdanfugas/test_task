package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/users", handleUsers)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var users []User
	json.Unmarshal(byteValue, &users)
	json.NewEncoder(w).Encode(users)
	defer jsonFile.Close()

	fmt.Println("Endpoint Hit: /users.")
}

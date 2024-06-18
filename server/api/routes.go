package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
	fmt.Println("got / routes")
}

func GetList(w http.ResponseWriter, r *http.Request) {

	array := [...]string{"1", "2"}

	jsonData, err := json.Marshal(array)

	if err != nil {
		http.Error(w, "Error converting json data", http.StatusInternalServerError)
		fmt.Println("Error /lists response: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Printf("got /lists route\n")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Users!")
	fmt.Println("got /users routes")
}

func StartServer() {
	http.HandleFunc("/", GetRoute)
	http.HandleFunc("/lists", GetList)
	http.HandleFunc("/users", GetUsers)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

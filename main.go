package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	port = "8080"
)

func main() {

	router := mux.NewRouter()
	router.Use()

	fmt.Printf("Start...")
	//начало работы сервера
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}

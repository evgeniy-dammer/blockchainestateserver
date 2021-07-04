package main

import (
	"blockchainestateserver/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mine", controllers.Mine).Methods("GET")
	router.HandleFunc("/transactions/new", controllers.NewTransaction).Methods("POST")
	router.HandleFunc("/chain", controllers.FullChain).Methods("GET")
	router.HandleFunc("/nodes/register", controllers.RegisterNodes).Methods("POST")
	router.HandleFunc("/nodes/resolve", controllers.Consensus).Methods("GET")

	fmt.Println("Blockchain is started...")
	http.ListenAndServe(":8181", router)
}

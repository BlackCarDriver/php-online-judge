package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getproblem", controllers.GetProblem)
	mux.HandleFunc("/commit", controllers.Commit)
	mux.HandleFunc("/gethistory", controllers.GetHistory)
	mux.HandleFunc("/getproblemlist", controllers.GetProblemList)
	fmt.Println("the server is listen on : 8083....")
	err := http.ListenAndServe("0.0.0.0:8083", mux)
	if err != nil {
		log.Fatal(err)
	}
}

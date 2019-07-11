package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
	"./models"
)

func main() {
	models.DBinit()
	mux := http.NewServeMux()

	// mux.HandleFunc("/example", errorHander(example))
	mux.HandleFunc("/getproblem", controllers.ErrorHander(controllers.GetProblem))
	mux.HandleFunc("/commit", controllers.ErrorHander(controllers.Commit))
	mux.HandleFunc("/gethistory", controllers.ErrorHander(controllers.GetHistory))
	fmt.Println("http服务器启动，端口：8083")
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal(err)
	}
}

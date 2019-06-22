package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// phpOJ.GenerateProject1Code()
	// phpOJ.RunProject1()
	// b, err := phpOJ.CheckProject1Answer()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(b)

	// models.InitDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/example", errorHander(example))

	fmt.Println("http服务器启动，端口：8083")
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("==")
}

func example(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	//data := getBodyData(r)
	fmt.Println("is Success")
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	w.Header().Set("content-type", "application/json")                           //返回数据格式是json
}

func getBodyData(r *http.Request) (data map[string]interface{}) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	err = json.Unmarshal(body, &data)
	checkErr(err)
	return
}

func errorHander(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), 500)
			}
		}()
		h(w, r)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

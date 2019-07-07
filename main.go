package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"
	"./phpOJ"
)

func main() {
	Server()
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
		}
	}()
	phpOJ.GenerateProject1Code()
	phpOJ.RunProject1()
	// b := phpOJ.CheckProject1Answer()
	// fmt.Println(b)

	// models.InitDB()
	// mux := http.NewServeMux()

	// mux.HandleFunc("/example", errorHander(example))

	// fmt.Println("http服务器启动，端口：8083")
	// err := http.ListenAndServe(":8083", mux)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("==")
}

func example(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	data := getBodyData(r)
	if userName, ok := data["userName"].(string); ok {
		fmt.Println(userName)

		phpOJ.GenerateProject1Code()
		phpOJ.RunProject1()
		b := phpOJ.CheckProject1Answer()
		fmt.Println(b)
	} else {
		panic("type assertion has error")
	}
	// fmt.Println("is Success")

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


//================================================================ 7-7

func Server(){
	mux := http.NewServeMux()
	mux.HandleFunc("/getproblem", GetProblem)
	mux.HandleFunc("/commit", Commit)

	server := &http.Server{
		Addr:           "localhost:6666",
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Begin to Listen!!")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}

//http://localhost:6666/getproblem
//provide the problem's data to defferent user and problem
func GetProblem(w http.ResponseWriter, r *http.Request){
	SetHeader(w)
	body, _ := ioutil.ReadAll(r.Body)
	if len(body)==0 {
		return
	}
	ssmap := getBodyMap(body)
	problemid := ssmap["pid"]
	userid := ssmap["uid"]
	tp := phpOJ.Problem{}
	if problemid=="" || userid=="" {
		WriteJson(w, tp)
	}
	tp = phpOJ.GetProblem(userid)
	WriteJson(w, tp)
}

//http://localhost:6666/commit
//pull or update user's code from github and then judge and return the result
func Commit(w http.ResponseWriter, r *http.Request){
	SetHeader(w)
	body, _ := ioutil.ReadAll(r.Body)
	if len(body)==0 {
		return
	}
	ssmap := getBodyMap(body)
	fmt.Println(ssmap)
	//do something ....
	WriteJson(w,"Accept")
}

func WriteJson(w http.ResponseWriter, data interface{}) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsondata)
}

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}

//parse request.Body to an string-string map
func getBodyMap(body []byte) map[string]string {
	var postbody map[string]string
	json.Unmarshal(body, &postbody)
	return postbody
}

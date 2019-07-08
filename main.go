package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./models"
	"./phpOJ"
)

func main() {
	models.DBinit()
	mux := http.NewServeMux()

	// mux.HandleFunc("/example", errorHander(example))
	mux.HandleFunc("/getproblem", errorHander(GetProblem))
	mux.HandleFunc("/commit", errorHander(Commit))
	fmt.Println("http服务器启动，端口：8083")
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal(err)
	}
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
				fmt.Println(err)
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

//http://localhost:6666/getproblem
//provide the problem's data to defferent user and problem
func GetProblem(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	// body, _ := ioutil.ReadAll(r.Body)
	// if len(body) == 0 {
	// 	return
	// }
	// ssmap := getBodyMap(body)
	// problemid := ssmap["pid"]
	// userid := ssmap["uid"]
	body := getBodyData(r)
	// uid, ok := body["uid"].(string)
	pid, ok := body["pid"].(float64)
	if !ok {
		err := fmt.Errorf("type assertion has error")
		panic(err)
	}

	tp := models.SelectProblem(int(pid))
	fmt.Println(tp)
	WriteJson(w, tp)
}

//http://localhost:6666/commit
//pull or update user's code from github and then judge and return the result
func Commit(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	uid, ok := body["uid"].(string)
	pid, ok := body["pid"].(float64)
	if !ok {
		err := fmt.Errorf("type assertion has error")
		panic(err)
	}
	tp := models.SelectProblem(int(pid))
	tu := models.SelectUser(uid)
	// fmt.Println(tp)
	// fmt.Println(tu)
	phpOJ.GitPull(tu.Repository, tu.Openid)
	phpOJ.GenerateProject1Code(tu.Openid, tp.CheckoutPath)
	defer phpOJ.GitCheckOut(tu.Openid)
	WriteJson(w, "Accept")
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

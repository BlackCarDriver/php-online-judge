package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../models"
)

//http://..../getproblem
//provide the problem's data to defferent user and problem
func GetProblem(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	pid, ok1 := body["pid"].(float64) //use float type to receive typescript number type
	uid, ok2 := body["uid"].(string)
	if !ok1 || !ok2 {
		return
		//panic("GetProblem() can not find pid and uid in require body...")
	}
	tp, err := OJS[int(pid)].GetProblem(int(pid), uid)
	checkErr(err)
	WriteJson(w, tp)
}

//http://...../commit
//pull or update user's code from github and then judge and return the result
func Commit(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	pid, ok1 := body["pid"].(float64)
	uid, ok2 := body["uid"].(string)
	if !ok1 || !ok2 {
		// panic("Commit() can not find pid and uid in require body...")
		return
	}
	res, err := OJS[int(pid)].Commit(int(pid), uid)
	checkErr(err)
	WriteJson(w, res)
}

//http://...../gethistory
//get and return user answer history
func GetHistory(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	pid, ok1 := body["pid"].(float64)
	uid, ok2 := body["uid"].(string)
	if !ok1 || !ok2 {
		// panic("GetHistory can not find pid and uid in require body...")
		return
	}
	res, err := models.GetHistory(int(pid), uid)
	checkErr(err)
	WriteJson(w, res)
}

//http://...../getproblemlist
//return the list of problem
func GetProblemList(w http.ResponseWriter, r *http.Request) {
	list := make([]string, 0)
	//do something...
	WriteJson(w, list)
}

//================================ tool function ====================

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

func WriteJson(w http.ResponseWriter, data interface{}) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsondata)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

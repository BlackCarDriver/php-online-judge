package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	"../models"
	"../phpOJ"
)

//http://..../getproblem
//provide the problem's data to defferent user and problem
func GetProblem(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	pid, ok := body["pid"].(float64)
	uid, ok2 := body["uid"].(string)
	if !ok || !ok2 {
		panic("can not find pid and uid in require body...")
	}
	tp := models.SelectProblem(int(pid))
	targeturl := phpOJ.PHPSubject1GetUrl(uid)
	giturl := phpOJ.GetUerGitUrl(pid, uid)
	tp.Description = fmt.Sprintf(tp.Description, targeturl, giturl)
	WriteJson(w, tp)
}

//http://...../commit
//pull or update user's code from github and then judge and return the result
func Commit(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	uid, ok1 := body["uid"].(string)
	pid, ok2 := body["pid"].(float64)
	if !ok1 || !ok2 {
		panic("can not find pid and uid in require body...")
	}
	tp := models.SelectProblem(int(pid))
	tu := models.SelectUser(uid)
	//phpOJ.GitPull(tu.Repository, tu.Openid)
	phpOJ.GenerateProject1Code(tu.Openid, tp.CheckoutPath)
	os.Exit(1)
	// defer phpOJ.GitCheckOut(tu.Openid)
	result := phpOJ.RunProject1(tu.Openid, tp.CheckoutPath)
	b := phpOJ.CheckProject1Answer(result)
	WriteJson(w, b)
}

//http://...../gethistory
//get and return user answer history
func GetHistory(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	body := getBodyData(r)
	pid, ok1 := body["pid"].(float64)
	uid, ok2 := body["uid"].(string)
	if !ok1 || !ok2 {
		panic("can not find pid and uid in require body...")
	}
	history := models.GetUserHistory(pid, uid)
	WriteJson(w, history)
}

//============== tool function ====================

func ErrorHander(h http.HandlerFunc) http.HandlerFunc {
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

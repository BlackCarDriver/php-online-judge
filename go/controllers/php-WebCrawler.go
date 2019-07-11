package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"../models"
)

type WebCrawler struct {
	OnlineJudge
}

var php01 = new(WebCrawler)

func (this *WebCrawler) GetProblem(pid int, uid string) (data models.Problem, err error) {
	problemData := models.SelectProblem(pid)
	targeturl := models.PHPSubject1GetUrl(uid)
	giturl := models.GetUerGitUrl(pid, uid)
	problemData.Description = fmt.Sprintf(problemData.Description, targeturl, giturl)
	return problemData, nil
}

func (this *WebCrawler) Commit(pid int, uid string) (result models.Result, err error) {
	tp := models.SelectProblem(pid)
	tu := models.SelectUser(uid)
	GitPull(tu.Repository, tu.Openid)
	fmt.Println("git pull scuess!!!")
	models.GenerateProject1Code(tu.Openid, tp.CheckoutPath)
	tresult := RunProject1(tu.Openid, tp.CheckoutPath)
	b := CheckProject1Answer(tresult)
	if b {
		result.Describe = "scuess!"
	} else {
		result.Describe = "fall!"
	}
	return result, nil
}

//=======================================================

//data return from docker container
type judgeResult struct {
	UserResult   string   `json:"userResult"`
	SystemResult []string `json:"systemResult"`
}

//bin/sh -c 后面记得用单引号，巨坑
// 最终输出是一个json格式
const (
	dockerRun = `
	PHP=$(pwd)
	sudo docker run \
		--rm \
		-v $PHP$0:/Code \
		php:alpine \
		/bin/sh -c '\

			userResult=$(php /Code/subject1.php); \
			result=$(echo $userResult | grep "error:"); \
			if [ $result != "" ];then \
				echo -n $userResult; \
			else \
				echo $userResult; \
				sysResult=$(php /Code/SystemCode.php); \
				echo -n $sysResult; \
			fi \

		';
	`
)

func RunProject1(openid string, checkout_path string) (result judgeResult) {
	codeUrl := fmt.Sprintf("/userData/%s%s", openid, checkout_path)
	params := make([]string, 3)
	params[0] = "-c"
	params[1] = dockerRun
	params[2] = codeUrl
	// re是一个json格式的结果
	r, err := execCommand("bash", params)
	// fmt.Println(r)
	checkErr(err)
	re := strings.Split(r, "\n")
	// fmt.Println(len(re))
	result.UserResult = re[0]
	json.Unmarshal([]byte(re[1]), &result.SystemResult)
	//将json转化为结构体
	// err = json.Unmarshal([]byte(re), &result)
	// checkErr(err)
	// fmt.Println(result)
	return
}

func CheckProject1Answer(result judgeResult) (b bool) {
	for _, v := range result.SystemResult {
		if !strings.Contains(result.UserResult, v) {
			b = false
			return
		}
	}
	b = true
	return
}

func execCommand(commandName string, params []string) (result string, err error) {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	//fmt.Println(cmd.Args)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		result = fmt.Sprint(err) + ": " + stderr.String()
	}
	result = out.String()
	return
}

func GitCheckOut(openid string) {
	args := make([]string, 2)
	args[0] = "-c"
	args[1] = fmt.Sprintf("cd ./userData/%s && git checkout .", openid)
	result, err := execCommand("/bin/bash", args)
	checkErr(err)
	fmt.Println(result)
}

func GitPull(gitUrl string, openid string) {
	args := make([]string, 2)
	args[0] = "-c"
	args[1] = fmt.Sprintf("cd ./userCode/%s && git pull %s", openid, gitUrl)
	result, err := execCommand("/bin/bash", args)
	checkErr(err)
	fmt.Println(result)
}

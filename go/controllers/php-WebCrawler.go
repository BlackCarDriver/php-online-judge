package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println("bging to git pull")
	GitPull(tu.Repository, tu.Openid)
	fmt.Println("begin to make project1 code")
	models.GenerateProject1Code(tu.Openid, tp.CheckoutPath)
	fmt.Println("begin to run project1")
	tresult := RunProject1(tu.Openid, tp.CheckoutPath)
	fmt.Println("begin to check project1")
	b := CheckProject1Answer(tresult)
	result.Time = "2019-11-11 (mock-data)"
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
		-v $0:/Code \
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
	codeUrl := fmt.Sprintf("%s/%s%s", models.UserCodePath, openid, checkout_path)
	fmt.Println("codeUrl : ", codeUrl)
	params := make([]string, 3)
	params[0] = "-c"
	params[1] = dockerRun
	params[2] = codeUrl
	// re是一个json格式的结果
	r, err := execCommand("bash", params)
	checkErr(err)
	re := strings.Split(r, "\n")
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
	args[1] = fmt.Sprintf("cd %s/%s && git checkout .", models.UserCodePath, openid)
	result, err := execCommand("/bin/bash", args)
	// checkErr(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func GitPull(gitUrl string, openid string) {
	fmt.Println("openid: ", openid, "   git url :", gitUrl)
	args := make([]string, 2)
	args[0] = "-c"
	args[1] = fmt.Sprintf("cd %s/%s && git pull --rebase %s", models.UserCodePath, openid, gitUrl)
	result, err := execCommand("/bin/bash", args)
	// checkErr(err)
	if err != nil {
		fmt.Println("GitPull() fall : ", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

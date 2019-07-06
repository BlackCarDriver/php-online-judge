package phpOJ

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

//bin/sh -c 后面记得用单引号，巨坑
// 最终输出是一个json格式
const (
	dockerRun = `
	echo -n '{'
	PHP=$(pwd)
	sudo docker run \
		--rm \
		-v $PHP/UserCode:/UserCode \
		-v $PHP/phpOJ/subject-1/SysTmpCode:/SysTmpCode \
		php:alpine \
		/bin/sh -c '\

			userResult=$(php /UserCode/zzm/test-UserCode.php); \
			result=$(echo $userResult | grep "error:"); \
			if [ $result != "" ];then \
				echo -n "\"userResult\":\""; \
				echo -n $userResult; \
				echo -n "\","; \
			else \
				echo -n "\"userResult\":\""; \
				echo -n $userResult; \
				echo -n "\","; \
				echo -n "\"systemResult\":"; \
				sysResult=$(php /SysTmpCode/SystemCode.php); \
				echo -n $sysResult; \
			fi \

		'; \
	echo -n '}';
	`
)

type Result struct {
	UserResult   string   `json:"userResult"`
	SystemResult []string `json:"systemResult"`
}

func RunProject1() {
	params := make([]string, 2)
	params[0] = "-c"
	params[1] = dockerRun
	// re是一个json格式的结果
	re, err := execCommand("bash", params)
	checkErr(err)
	var result Result
	//将json转化为结构体
	err = json.Unmarshal([]byte(re), &result)
	checkErr(err)
	fmt.Println(result)
}

func GenerateProject1Code() {
	phpfile, err := os.Create("./phpOJ/subject-1/SysTmpCode/SystemCode.php")
	checkErr(err)
	defer phpfile.Close()
	url := "https://blog.csdn.net/YDTG1993/article/details/83861629"
	tmpl, err := template.ParseFiles("./phpOJ/subject-1/SysTmpCode/php-template.txt")
	checkErr(err)
	err = tmpl.Execute(phpfile, url)
	checkErr(err)
}

func CheckProject1Answer() (b bool) {

	sysResult, err := ioutil.ReadFile("./phpOJ/subject-1/SysTmpCode/SystemResult.txt")
	checkErr(err)
	userResult, err := ioutil.ReadFile("./phpOJ/subject-1/SysTmpCode/zzm/UserResult.txt")
	checkErr(err)
	sys := strings.Split(string(sysResult), "\n")
	user := string(userResult)
	for _, v := range sys {
		if !strings.Contains(user, v) {
			b = false
			return
		}
	}
	b = true
	return
}

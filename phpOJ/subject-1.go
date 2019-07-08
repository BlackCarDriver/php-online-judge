package phpOJ

import (
	"encoding/json"
	"fmt"
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
		-v $PHP$0:/Code \
		php:alpine \
		/bin/sh -c '\

			userResult=$(php /Code/subject1.php); \
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
				sysResult=$(php /Code/SystemCode.php); \
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

func RunProject1(openid string, checkout_path string) (result Result) {
	codeUrl := fmt.Sprintf("/userData/%s%s", openid, checkout_path)
	params := make([]string, 3)
	params[0] = "-c"
	params[1] = dockerRun
	params[2] = codeUrl
	// re是一个json格式的结果
	re, err := execCommand("bash", params)
	checkErr(err)
	//将json转化为结构体
	err = json.Unmarshal([]byte(re), &result)
	checkErr(err)
	fmt.Println(result)
	return
}

func GenerateProject1Code(openid string, checkout_path string) {
	codeUrl := fmt.Sprintf("./userData/%s%s", openid, checkout_path)
	phpfile, err := os.Create(codeUrl + "SystemCode.php")
	checkErr(err)
	defer phpfile.Close()
	url := "https://blog.csdn.net/YDTG1993/article/details/83861629"
	tmpl, err := template.ParseFiles("./phpOJ/subject-1/SysTmpCode/php-template.txt")
	checkErr(err)
	err = tmpl.Execute(phpfile, url)
	checkErr(err)
}

func CheckProject1Answer(result Result) (b bool) {
	for _, v := range result.SystemResult {
		if !strings.Contains(result.UserResult, v) {
			b = false
			return
		}
	}
	b = true
	return
}

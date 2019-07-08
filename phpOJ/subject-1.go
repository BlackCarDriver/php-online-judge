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

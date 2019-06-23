package phpOJ

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const (
	dockerRun = `
	PHP=$(pwd)
	sudo docker run \
		--rm \
		-i \
		-v $PHP/UserCode:/UserCode \
		-v $PHP/phpOJ/subject-1/SysTmpCode:/SysTmpCode \
		php:alpine \
		/bin/sh -c "\

			userResult=$(php /UserCode/zzm/test-UserCode.php); \
			result=$(echo "$userResult" | grep "error:"); \
			if [ "$result" != "" ];then \
				echo $userResult; \
			else \
				echo "$userResult" > /SysTmpCode/zzm/UserResult.txt; \
				sysResult=$(php /SysTmpCode/SystemCode.php); \
				echo "$sysResult" > /SysTmpCode/SystemResult.txt; \
			fi \

		"
	`
)

func RunProject1() {
	// os.Chdir("../shell")
	params := make([]string, 2)
	params[0] = "-c"
	params[1] = dockerRun
	// params[1] = "php.sh"
	err := execCommand("bash", params)
	checkErr(err)
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

package phpOJ

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func RunProject1() {
	// os.Chdir("../shell")
	params := make([]string, 1)
	params[0] = "./phpOJ/subject-1/SysTmpCode/dockerRunPHP.sh"
	// params[1] = "php.sh"
	execCommand("bash", params)
}

func GenerateProject1Code() {
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
		}
	}()
	path, err := os.Getwd()
	checkErr(err)
	err = os.Chdir("./phpOJ/subject-1/SysTmpCode")
	checkErr(err)
	phpfile, err := os.Create("SystemCode.php")
	checkErr(err)
	url := "https://blog.csdn.net/YDTG1993/article/details/83861629"
	tmpl, err := template.ParseFiles("php-template.txt")
	checkErr(err)
	err = tmpl.Execute(phpfile, url)
	checkErr(err)
	err = os.Chdir(path)
	checkErr(err)
}

func CheckProject1Answer() (b bool, err error) {

	sysResult, err := ioutil.ReadFile("./phpOJ/subject-1/SysTmpCode/SystemResult.txt")
	if err != nil {
		return
	}
	userResult, err := ioutil.ReadFile("./phpOJ/subject-1/SysTmpCode/zzm/UserResult.txt")
	if err != nil {
		return
	}
	sys := strings.Split(string(sysResult), "\n")
	user := string(userResult)
	for _, v := range sys {
		if !strings.Contains(user, v) {
			b = false
		}
	}
	b = true
	return
}

package phpOJ

import (
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
		}
	}
	b = true
	return
}

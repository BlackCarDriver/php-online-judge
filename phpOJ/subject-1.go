package phpOJ

import (
	"fmt"
	"os"
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

func CheckProject1Answer() bool {
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
		}
	}()

}

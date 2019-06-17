package phpOJ

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const (
	shellPath  = "./shell"
	configPath = "./phpOJ/conf/"
)

var (
	phpConf        ConfigMachine
	urlList        []string
	urlListSize    int
	probemTemplate string
)

func init() {
	var err error
	phpConf, err = NewConfig(configPath)
	handleErr("NewConfig(configPath)", err, true)
	//get url list from config file
	err = phpConf.Register("url_list", make([]string, 0), true)
	if err != nil {
		fmt.Println(err)
	} else {
		url, _ := phpConf.Get("url_list")
		urlList = url.([]string)
		urlListSize = len(urlList)
		if urlListSize == 0 {
			log.Fatal("urlList config unright!")
		}
	}
	//get problem template from config file
	err = phpConf.Register("problem_template", "", true)
	if err != nil {
		log.Fatal(err)
	} else {
		tmp, _ := phpConf.Get("problem_template")
		probemTemplate = tmp.(string)
	}

	fmt.Println(getProblemText("1234567"))
}

// the entrance of it package
func Main() {
	// os.Chdir("../shell")
	params := make([]string, 1)
	params[0] = "./phpOJ/shell/dockerRunPHP.sh"
	// params[1] = "php.sh"
	execCommand("bash", params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	//fmt.Println(cmd.Args)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	fmt.Print(out.String())
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

package phpOJ

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	shellPath  = "./shell"
	configPath = "./phpOJ/conf/"
)

// var (
// 	phpConf        ConfigMachine
// 	urlList        []string
// 	urlListSize    int
// 	probemTemplate string
// 	gitUrlTemplate string
// 	userCodePath   string
// )

// func init() {
// 	//init config values directly
// 	userCodePath = `./UserCode`

// 	var err error
// 	phpConf, err = NewConfig(configPath)
// 	handleErr("NewConfig(configPath)", err, true)
// 	//get url list from config file
// 	err = phpConf.Register("url_list", make([]string, 0), true)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		url, _ := phpConf.Get("url_list")
// 		urlList = url.([]string)
// 		urlListSize = len(urlList)
// 		if urlListSize == 0 {
// 			log.Fatal("urlList config unright!")
// 		}
// 	}
// 	//get problem template from config file
// 	err = phpConf.Register("problem_template", "", true)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		tmp, _ := phpConf.Get("problem_template")
// 		probemTemplate = tmp.(string)
// 	}
// 	//get github url template from config file
// 	err = phpConf.Register("gitUrl_template", "", true)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		tmp, _ := phpConf.Get("gitUrl_template")
// 		gitUrlTemplate = tmp.(string)
// 	}

// 	// ================ the following is test code ==============

// 	//fmt.Println(getProblemText("UserName"))
// }

// the entrance of it package
// func Main() {
// 	// os.Chdir("../shell")
// 	params := make([]string, 1)
// 	params[0] = "./phpOJ/shell/dockerRunPHP.sh"
// 	// params[1] = "php.sh"
// 	execCommand("bash", params)
// }

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

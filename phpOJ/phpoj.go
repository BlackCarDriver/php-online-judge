package phpOJ

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	shellPath    = "./shell"
	userCodePath = ""
)

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

func GitPull(gitUrl string, openid string) {
	args := make([]string, 2)
	args[0] = "-c"
	args[1] = fmt.Sprintf("cd ./userData/%s && git pull %s", openid, gitUrl)
	result, err := execCommand("/bin/bash", args)
	checkErr(err)
	fmt.Println(result)
}

func GitCheckOut(openid string) {
	args := make([]string, 2)
	args[0] = "-c"
	args[1] = fmt.Sprintf("cd ./userData/%s && git checkout .", openid)
	result, err := execCommand("/bin/bash", args)
	checkErr(err)
	fmt.Println(result)
}

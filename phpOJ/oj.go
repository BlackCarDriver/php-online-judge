package oj

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	shellPath = "./shell"
)

func Run() {
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

//the entrance of it package
func Main() {
	Run()
}

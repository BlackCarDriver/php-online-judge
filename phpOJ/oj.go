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
	params := make([]string, 2)
	params[0] = "-c"
	params[1] = `
	PHP=$(pwd)
	docker run --name php \
		--rm \
		-i \
		-v $PHP/code:/code \
		php \
		php /code/index.php`
	// params[1] = "php.sh"
	execCommand("bash", params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	fmt.Println(out.String())
	return true
}

//the entrance of it package
func Main() {
	Run()
}

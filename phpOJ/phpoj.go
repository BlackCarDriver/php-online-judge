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

var (
	phpConf ConfigMachine
)

func init() {
	var err error
	phpConf, err = NewConfig(configPath)
	handleErr("NewConfig(configPath)", err, true)
	err = phpConf.Register("test", "", true)
	if handleErr("Register", err, false) == false {
		phpConf.Display()
	}
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

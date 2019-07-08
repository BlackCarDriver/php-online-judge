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

// var (
// 	urlList        []string
// 	probemTemplate string
// 	gitUrlTemplate string
// )

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

//======================================== 7-7

// func init() {
// 	urlList = []string{
// 		"https://studygolang.com/pkgdoc",
// 		"https://www.runoob.com/",
// 		"https://www.jb51.net/list/list_172_1.htm",
// 		"http://www.fhdq.net/emoji/emojifuhao.html",
// 		"https://www.oschina.net/",
// 	}
// 	probemTemplate = `
// 	Here we create an array a that will hold exactly 5 ints.
// 	The type of elements and length are both part of the array’s type.
// 	By default an array is zero-valued, which for ints means 0s.
// 	Your target is %s ..... and you need to pull your code to %s !`

// 	gitUrlTemplate = "https://github.com/BlackCarDriver/%s.git"
// }

// //the message of a problem
// type Problem struct {
// 	Text  string `json:"text"`
// 	Time  string `json:"time"`
// 	Type  string `json:"type"`
// 	Tag   string `json:"tag"`
// 	Try   int    `json:"try"`
// 	Ac    int    `json:"ac"`
// 	Rate  string `json:"rate"`
// 	Leave string `json:"leave"`
// }

// //distribute different url of target website to user by uesr's id
// func getUrlById(id string) string {
// 	idHash := int(crc32.ChecksumIEEE([]byte(id)))
// 	if idHash < 0 {
// 		idHash = -idHash
// 	}
// 	return urlList[idHash%len(urlList)]
// }

// //get url of user's github
// func getGitUrlById(id string) string {
// 	return fmt.Sprintf(gitUrlTemplate, id)
// }

// //create the problem message according to userid
// func GetProblem(id string) Problem {
// 	problem_url := getUrlById(id)
// 	github_url := getGitUrlById(id)
// 	//the following message should get from database
// 	temp := Problem{Text: "", Time: "2019-6-6", Type: "PHP", Tag: "网络爬虫", Try: 133, Ac: 40, Leave: "⭐⭐⭐⭐"}
// 	temp.Rate = fmt.Sprintf("%%%.1f", float32(temp.Ac*100/temp.Try))
// 	temp.Text = fmt.Sprintf(probemTemplate, problem_url, github_url)
// 	temp.Text = strings.Replace(temp.Text, "\n", "<br>", -1)
// 	return temp
// }

//pull user code
// func updataCodeById(id string) error {
// 	gitUrl := getGitUrlById(id)
// 	args := make([]string, 3)
// 	args[0] = "-c"
// 	args[1] = fmt.Sprintf(`cd %s && git clone %s`, userCodePath, gitUrl)
// 	//err := execCommand("bash", args)
// 	fmt.Println("user code already pulled!....")
// 	return nil
// }

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

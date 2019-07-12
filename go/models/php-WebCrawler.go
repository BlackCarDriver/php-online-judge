package models

import (
	"database/sql"
	"fmt"
	"hash/crc32"
	"os"
	"text/template"
	"time"
)

const (
	UserCodePath = `/home/ubuntu/Desktop/userCode`
)

//list of target url in subject-1
var PHPSubject1Url = []string{
	"https://blog.csdn.net/YDTG1993/article/details/83861629",
	"https://studygolang.com/articles/16010?fr=sidebar",
	"https://www.csdn.net",
	"https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin",
	"https://beego.me/docs/intro",
	"https://gocn.vip",
}

//distribut target url randomly by user id
func PHPSubject1GetUrl(id string) string {
	idHash := int(crc32.ChecksumIEEE([]byte(id)))
	if idHash < 0 {
		idHash = -idHash
	}
	return PHPSubject1Url[idHash%len(PHPSubject1Url)]
}

func GenerateProject1Code(openid string, checkout_path string) {
	codeUrl := fmt.Sprintf("%s/%s%s", UserCodePath, openid, checkout_path)
	phpfile, err := os.Create(codeUrl + "SystemCode.php")
	checkErr(err)
	defer phpfile.Close()
	url := PHPSubject1GetUrl(openid)
	tmpl, err := template.ParseFiles("./answer/webCrawler/answer-template.txt")
	checkErr(err)
	err = tmpl.Execute(phpfile, url)
	checkErr(err)
}

func SelectProblem(pid int) (t Problem) {
	rows, err := db.Query("select id,title,description,type,code_type,checkout_path,attach_code,attach_file,answer,status,create_time from t_problem where id = $1", pid)
	checkErr(err)
	defer rows.Close()
	if rows.Next() {
		var Id sql.NullInt64
		var Title sql.NullString
		var Description sql.NullString
		var Type sql.NullString
		var CodeType sql.NullString
		var CheckoutPath sql.NullString
		var AttachCode sql.NullString
		var AttachFile sql.NullString
		var Answer sql.NullString
		var Status bool
		var CreateTime time.Time
		err := rows.Scan(&Id, &Title, &Description, &Type, &CodeType, &CheckoutPath, &AttachCode, &AttachFile, &Answer, &Status, &CreateTime)
		checkErr(err)
		t.Id = Id.Int64
		t.Title = Title.String
		t.Description = Description.String
		t.Type = Type.String
		t.CodeType = CodeType.String
		t.CheckoutPath = CheckoutPath.String
		t.AttachCode = AttachCode.String
		t.AttachFile = AttachFile.String
		t.Answer = Answer.String
		t.Status = Status
		t.CreateTime = CreateTime.Format("2006-01-02")
	}
	return
}

func SelectUser(uid string) (t TUser) {
	rows, err := db.Query("select openid,nickname,avatar_url,git_account,git_password,repository,create_time from t_user where openid=$1", uid)
	checkErr(err)
	defer rows.Close()
	if rows.Next() {
		var Openid sql.NullString
		var Nickname sql.NullString
		var AvatarUrl sql.NullString
		var GitAccount sql.NullString
		var GitPassword sql.NullString
		var Repository sql.NullString
		var CreateTime time.Time
		err := rows.Scan(&Openid, &Nickname, &AvatarUrl, &GitAccount, &GitPassword, &Repository, &CreateTime)
		checkErr(err)
		t.Openid = Openid.String
		t.Nickname = Nickname.String
		t.AvatarUrl = AvatarUrl.String
		t.GitAccount = GitAccount.String
		t.GitPassword = GitPassword.String
		t.Repository = Repository.String
		t.CreateTime = CreateTime.Format("2006-01-02")
	}
	return
}

//return the url which user should upload their answer to
func GetUerGitUrl(problemid int, userid string) string {
	//do something...
	return "https://github.com/gopher/upload  (mock data)"
}

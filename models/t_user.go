package models

import (
	"time"
)

type TUser struct {
	Openid      string    `json:"openid"`
	Nickname    string    `json:"nickname"`
	AvatarUrl   string    `json:"avatar_url"`
	GitAccount  string    `json:"git_account"`
	GitPassword string    `json:"git_password"`
	Repository  string    `json:"repository"`
	CreateTime  time.Time `json:"create_time"`
}

// func (t *TUser) Insert() {
// 	stmt, err := db.Prepare("insert into t_user(openid,nickname,avatar_url,git_account,git_password,repository,create_time) values ($1,$2,$3,$4,$5,$6,$7)")
// 	checkErr(err)
// 	_, err = stmt.Exec(t.Openid, t.Nickname, t.AvatarUrl, t.GitAccount, t.GitPassword, t.Repository, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TUser) Update() {
// 	stmt, err := db.Prepare("update t_user set openid=$1,nickname=$2,avatar_url=$3,git_account=$4,git_password=$5,repository=$6,create_time=$7") //where语句自行定义
// 	checkErr(err)
// 	_, err = stmt.Exec(t.Openid, t.Nickname, t.AvatarUrl, t.GitAccount, t.GitPassword, t.Repository, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TUser) Delete() {
// 	stmt, err := db.Prepare("delete from t_user") //where语句自行定义
// 	checkErr(err)
// 	_, err = stmt.Exec()
// 	checkErr(err)
// }
// func Select() (ts []TUser) {
// 	rows, err := db.Query("select * from t_user")
// 	checkErr(err)
// 	for rows.Next() {
// 		var t TUser
// 		var Openid sql.NullString
// 		var Nickname sql.NullString
// 		var AvatarUrl sql.NullString
// 		var GitAccount sql.NullString
// 		var GitPassword sql.NullString
// 		var Repository sql.NullString
// 		var CreateTime time.Time
// 		err := rows.Scan(&Openid, &Nickname, &AvatarUrl, &GitAccount, &GitPassword, &Repository, &CreateTime)
// 		checkErr(err)
// 		t.Openid = Openid.String
// 		t.Nickname = Nickname.String
// 		t.AvatarUrl = AvatarUrl.String
// 		t.GitAccount = GitAccount.String
// 		t.GitPassword = GitPassword.String
// 		t.Repository = Repository.String
// 		t.CreateTime = CreateTime
// 		ts = append(ts, t)
// 	}
// 	return
// }

package models

import (
	"time"
)

type TRecord struct {
	ProblemId    int64     `json:"problem_id"`
	Openid       string    `json:"openid"`
	CommitHash   string    `json:"commit_hash"`
	AnswerStatus bool      `json:"answer_status"`
	CreateTime   time.Time `json:"create_time"`
}

// func (t *TRecord) Insert() {
// 	stmt, err := db.Prepare("insert into t_record(problem_id,openid,commit_hash,answer_status,create_time) values ($1,$2,$3,$4,$5)")
// 	checkErr(err)
// 	_, err = stmt.Exec(t.ProblemId, t.Openid, t.CommitHash, t.AnswerStatus, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TRecord) Update() {
// 	stmt, err := db.Prepare("update t_record set problem_id=$1,openid=$2,commit_hash=$3,answer_status=$4,create_time=$5") //where语句自行定义
// 	checkErr(err)
// 	_, err = stmt.Exec(t.ProblemId, t.Openid, t.CommitHash, t.AnswerStatus, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TRecord) Delete() {
// 	stmt, err := db.Prepare("delete from t_record") //where语句自行定义
// 	checkErr(err)
// 	_, err = stmt.Exec()
// 	checkErr(err)
// }
// func Select() (ts []TRecord) {
// 	rows, err := db.Query("select * from t_record")
// 	checkErr(err)
// 	for rows.Next() {
// 		var t TRecord
// 		var ProblemId sql.NullInt64
// 		var Openid sql.NullString
// 		var CommitHash sql.NullString
// 		var AnswerStatus bool
// 		var CreateTime time.Time
// 		err := rows.Scan(&ProblemId, &Openid, &CommitHash, &AnswerStatus, &CreateTime)
// 		checkErr(err)
// 		t.ProblemId = ProblemId.Int64
// 		t.Openid = Openid.String
// 		t.CommitHash = CommitHash.String
// 		t.AnswerStatus = AnswerStatus
// 		t.CreateTime = CreateTime
// 		ts = append(ts, t)
// 	}
// 	return
// }

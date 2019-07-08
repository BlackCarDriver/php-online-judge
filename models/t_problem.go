package models

import (
	"database/sql"
	"time"
)

type TProblem struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	CodeType     string    `json:"code_type"`
	CheckoutPath string    `json:"checkout_path"`
	AttachCode   string    `json:"attach_code"`
	AttachFile   string    `json:"attach_file"`
	Answer       string    `json:"answer"`
	Status       bool      `json:"status"`
	CreateTime   time.Time `json:"create_time"`
}

// func (t *TProblem) Insert() {
// 	stmt, err := db.Prepare("insert into t_problem(id,title,description,type,code_type,checkout_path,attach_code,attach_file,answer,status,create_time) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
// 	checkErr(err)
// 	_, err = stmt.Exec(t.Id, t.Title, t.Description, t.Type, t.CodeType, t.CheckoutPath, t.AttachCode, t.AttachFile, t.Answer, t.Status, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TProblem) Update() {
// 	stmt, err := db.Prepare("update t_problem set id=$1,title=$2,description=$3,type=$4,code_type=$5,checkout_path=$6,attach_code=$7,attach_file=$8,answer=$9,status=$10,create_time=$11") //where语句自行定义
// 	defer stmt.Close()
// 	checkErr(err)
// 	_, err = stmt.Exec(t.Id, t.Title, t.Description, t.Type, t.CodeType, t.CheckoutPath, t.AttachCode, t.AttachFile, t.Answer, t.Status, t.CreateTime)
// 	checkErr(err)
// }
// func (t *TProblem) Delete() {
// 	stmt, err := db.Prepare("delete from t_problem") //where语句自行定义
// 	defer stmt.Close()
// 	checkErr(err)
// 	_, err = stmt.Exec()
// 	checkErr(err)
// }
func SelectProblem(pid int) (t TProblem) {
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
		t.CreateTime = CreateTime
	}
	return
}

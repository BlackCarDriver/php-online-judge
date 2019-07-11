package models

type Problem struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	CodeType     string `json:"code_type"`
	CheckoutPath string `json:"checkout_path"`
	AttachCode   string `json:"attach_code"`
	AttachFile   string `json:"attach_file"`
	Answer       string `json:"answer"`
	Status       bool   `json:"status"`
	CreateTime   string `json:"create_time"`
}

//result of judge
type Result struct {
	Time     string `json:"time"`
	State    int    `json:"state"`
	Describe string `json:"describe"`
}

type TUser struct {
	Openid      string `json:"openid"`
	Nickname    string `json:"nickname"`
	AvatarUrl   string `json:"avatar_url"`
	GitAccount  string `json:"git_account"`
	GitPassword string `json:"git_password"`
	Repository  string `json:"repository"`
	CreateTime  string `json:"create_time"`
}

//commit history
type Record struct {
	Create_time   string `json:"create_time"`
	Problem_id    int32  `json:"problem_id"`
	Openid        string `json:"openid"`
	Commit_hash   string `json:"commit_hash"`
	Answer_status bool   `json:"answer_status"`
}

//get user commit history and result of commit
func GetHistory(pid int, uid string) (history Record, err error) {
	//do something...
	history.Create_time = "2019-11-11"
	history.Commit_hash = "正确通过！  (mock data)"
	return
}

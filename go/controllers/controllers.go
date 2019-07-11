package controllers

import (
	"../models"
)

type OnlineJudge interface {
	GetProblem(pid int, uid string) (data models.Problem, err error)
	Commit(pid int, uid string) (result models.Result, err error)
}

var OJS = make(map[int]OnlineJudge, 0)

func Register(pid int, oj OnlineJudge) {
	if _, ok := OJS[pid]; ok {
		panic("This problem id have been used !")
	}
	OJS[pid] = oj
}

func init() {
	Register(1, php01)
}

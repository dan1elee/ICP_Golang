package idhandler

import (
	"github.com/google/uuid"
)

var courseIds map[string]bool = map[string]bool{}
var evalIds map[string]bool = map[string]bool{}

func GenCourseId() string {
	uid := uuid.New().String()
	_, exist := courseIds[uid]
	for exist {
		uid := uuid.New().String()
		_, exist = courseIds[uid]
	}
	return uid
}

func GenEvalId() string {
	uid := uuid.New().String()
	_, exist := evalIds[uid]
	for exist {
		uid := uuid.New().String()
		_, exist = evalIds[uid]
	}
	return uid
}

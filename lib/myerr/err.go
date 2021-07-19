package myerr

import (
	"path/filepath"
	"runtime"
	"strconv"
)

/*
自定义一个error对象，当我们链式调用函数时，错误也可以链式追踪，而不仅仅停留在表层。
*/

func New(text string) error {
	_, file, line, _ := runtime.Caller(1)
	filename := filepath.Base(file)
	return &errorString{
		filename: filename,
		line:     line,
		s:        text}
}

type errorString struct {
	filename string
	line     int
	s        string
}

func (e *errorString) Error() string {
	return e.filename + ":" + strconv.Itoa(e.line) + " " + e.s
}

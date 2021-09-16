package wraperrors

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// go 1.13推出的错误处理方案
// 自定义错误类型
type wrapError struct {
	Code   int    `json:"code"`   // 错误ID
	Detail string `json:"detail"` // 错误描述
	Err    error  `json:"-"`      // 实际错误
}

// 返回错误描述
func (e *wrapError) Error() string {
	txt, _ := json.Marshal(e)
	return string(txt)
}

func (e *wrapError) Unwrap() error {
	return e.Err
}

func (e *wrapError) Trace() string {
	var buff bytes.Buffer
	buff.WriteString("{")
	buff.WriteString(fmt.Sprintf("\"code\":%d,", e.Code))
	buff.WriteString(fmt.Sprintf("\"detail\":\"%s\"", e.Detail))
	if nil != e.Err {
		verr, ok := e.Err.(*wrapError)
		if !ok {
			buff.WriteString(fmt.Sprintf(",\"err\":\"%s\"", e.Err.Error()))
		} else {
			buff.WriteString(fmt.Sprintf(",\"err\":%s", verr.Trace()))
		}
	}
	buff.WriteString("}")
	return buff.String()
}

func equal(a, b error) bool {
	if nil == a || nil == b {
		return a == b
	}
	va, ok1 := a.(*wrapError)
	vb, ok2 := b.(*wrapError)
	if ok1 && ok2 && va.Code == vb.Code && va.Detail == vb.Detail {
		return true
	}
	return va.Error() == vb.Error()
}

func (e *wrapError) Is(err error) bool {
	if nil == err || nil == e {
		return e == err
	}
	if equal(e, err) {
		return true
	}
	return equal(e.Err, err)
}

func New(code int, detail string) *wrapError {
	return &wrapError{
		Code:   code,
		Detail: detail,
	}
}

func Wrap(code int, detail string, err error) *wrapError {
	return &wrapError{
		Code:   code,
		Detail: detail,
		Err:    err,
	}
}

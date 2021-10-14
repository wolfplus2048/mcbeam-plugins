package wraperrors

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// go 1.13推出的错误处理方案
// 自定义错误类型
type WrapError struct {
	Code   int    `json:"code"`   // 错误ID
	Detail string `json:"detail"` // 错误描述
	Err    error  `json:"-"`      // 实际错误
}

// 返回错误描述
func (e *WrapError) Error() string {
	txt, _ := json.Marshal(e)
	return string(txt)
}

func (e *WrapError) Unwrap() error {
	return e.Err
}

func (e *WrapError) Trace() string {
	var buff bytes.Buffer
	buff.WriteString("{")
	buff.WriteString(fmt.Sprintf("\"code\":%d,", e.Code))
	buff.WriteString(fmt.Sprintf("\"detail\":\"%s\"", e.Detail))
	if nil != e.Err {
		verr, ok := e.Err.(*WrapError)
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
	va, ok1 := a.(*WrapError)
	vb, ok2 := b.(*WrapError)
	if ok1 && ok2 && va.Code == vb.Code && va.Detail == vb.Detail {
		return true
	}
	return va.Error() == vb.Error()
}

func (e *WrapError) Is(err error) bool {
	if nil == err || nil == e {
		return e == err
	}
	if equal(e, err) {
		return true
	}
	return equal(e.Err, err)
}

// 创建一个wrapError
func New(code int, detail string) *WrapError {
	return &WrapError{
		Code:   code,
		Detail: detail,
	}
}

// 创建一个wrapError
func Wrap(code int, detail string, err error) *WrapError {
	return &WrapError{
		Code:   code,
		Detail: detail,
		Err:    err,
	}
}

// 返回wrapError格式字符串
func Errorf(code int, format string, a ...interface{}) string {
	e := WrapError{
		Code:   code,
		Detail: fmt.Sprintf(format, a...),
	}
	return e.Error()
}

// 解析错误字符
func Parse(msg string) (*WrapError, error) {
	e := WrapError{}
	err := json.Unmarshal([]byte(msg), &e)
	if nil != err {
		return nil, err
	}
	return &e, nil
}

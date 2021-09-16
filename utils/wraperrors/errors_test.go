package wraperrors

import (
	"errors"
	"fmt"
	"testing"
)

var Err01 = &wrapError{
	Code:   1001,
	Detail: "test 01",
}
var Err02 = &wrapError{
	Code:   1002,
	Detail: "test 02",
}
var Err03 = &wrapError{
	Code:   1003,
	Detail: "test 03",
}
var Err04 = errors.New("raw error")
var testData = []*wrapError{
	Err01,
	Err02,
	Err03,
}

func TestErrors(t *testing.T) {

	for _, ep := range testData {
		got := New(ep.Code, ep.Detail)
		if got.Error() != ep.Error() {
			t.Fatalf("Expected %s got %s", ep.Error(), got.Error())
		}

		if got.Code != ep.Code {
			t.Fatalf("Expected %d got %d", ep.Code, got.Code)
		}

		if got.Detail != ep.Detail {
			t.Fatalf("Expected %s got %s", ep.Detail, got.Detail)
		}

	}
}

func TestWrap(t *testing.T) {
	var got error
	for _, ep := range testData {
		if got == nil {
			got = ep
		} else {
			got = Wrap(ep.Code, ep.Detail, got)
		}
	}

	for _, it := range testData {
		if !errors.Is(got, it) {
			t.Fatalf("Is(%v, %v) want %v", got, it, it)
		}
	}

	var as *wrapError
	if !errors.As(got, &as) {
		t.Fatalf("Expected %v got %v", as, got)
	} else {
		t.Log(as)
	}
}

func TestTrace(t *testing.T) {
	ne := "{\"code\":1003,\"detail\":\"test 03\",\"err\":{\"code\":1002,\"detail\":\"test 02\",\"err\":{\"code\":1001,\"detail\":\"test 01\"}}}"
	var got error
	for _, ep := range testData {
		if got == nil {
			got = ep
		} else {
			got = Wrap(ep.Code, ep.Detail, got)
		}
	}
	got = fmt.Errorf("raw %w", got)

	//t.Log(got)
	var as *wrapError
	if !errors.As(got, &as) {
		t.Fatalf("Expected %v got %v", as, got)
	}

	if as.Trace() != ne {
		t.Fatalf("Expected %v got %v", ne, as.Trace())
	}
}

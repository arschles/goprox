package assert

import (
	"errors"
	"fmt"
	"testing"
)

type testCase struct {
	val  interface{}
	name string
}

var (
	nonNilTestCases = []testCase{
		testCase{val: 0, name: "0"},
		testCase{val: 1, name: "1"},
		testCase{val: true, name: "true"},
		testCase{val: false, name: "false"},
		testCase{val: "abc", name: `"abc"`},
		testCase{val: struct{}{}, name: "struct{}{}"},
	}
)

func TestCallerStr(t *testing.T) {
	s := callerStr(0)
	if len(s) <= 0 {
		t.Errorf("return value of callerStr is empty")
	}
}

func TestCallerStrf(t *testing.T) {
	fmtStr := "%d%d"
	val1 := 1
	val2 := 2
	res := callerStrf(0, fmtStr, val1, val2)
	if len(res) < 3 {
		t.Errorf("return value of callerStrf is not long enough")
	}
}

func TestTrue(t *testing.T) {
	rt := newRecordingTester()
	True(rt, true, "true wasn't reported as true")
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("no fatals expected, got [%d] instead", n)
	}
	True(rt, false, "false was reported as true")
	if n := rt.fatalsLen(); n != 1 {
		t.Fatalf("one fatal expected, got [%d] instead", n)
	}
}

func TestFalse(t *testing.T) {
	rt := newRecordingTester()
	False(rt, false, "false wasn't reported as false")
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("no fatals expected, got [%d] instead", n)
	}
	False(rt, true, "true wasn't reported as true")
	if n := rt.fatalsLen(); n != 1 {
		t.Fatalf("1 fatal expected, got [%d] instead", n)
	}
}

func TestNil(t *testing.T) {
	rt := newRecordingTester()
	var i *int = nil
	var b *bool = nil
	var s *string = nil
	var slc *[]string = nil
	var str *struct{} = nil
	tests := []testCase{
		testCase{val: nil, name: "nil"},
		testCase{val: i, name: "*int"},
		testCase{val: b, name: "*bool"},
		testCase{val: s, name: "*string"},
		testCase{val: slc, name: "*[]string"},
		testCase{val: str, name: "*struct{}"},
	}
	for i, test := range tests {
		Nil(rt, test.val, test.name)
		if n := rt.fatalsLen(); n != 0 {
			t.Fatalf("expected 0 fatals, got [%d] instead (iter [%d])", n, i)
		}
	}
}

func TestNotNil(t *testing.T) {
	rt := newRecordingTester()
	for i, test := range nonNilTestCases {
		NotNil(rt, test.val, test.name)
		if n := rt.fatalsLen(); n != 0 {
			t.Fatalf("expected 0 fatals, got [%d] instead (iter [%d])", n, i)
		}
	}
}

func TestErr(t *testing.T) {
	rt := newRecordingTester()
	err11 := errors.New("err1")
	err12 := errors.New("err1")
	Err(rt, err11, err12)
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("expected 0 fatals, got [%d] instead", n)
	}
	err21 := fmt.Errorf("err2-%s", "a")
	err22 := fmt.Errorf("err2-%s", "a")
	Err(rt, err21, err22)
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("expected 0 fatals, got [%d] instead", n)
	}
}

func TestExistsErr(t *testing.T) {
	rt := newRecordingTester()
	ExistsErr(rt, errors.New("abc"), "error")
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("expected 0 fatals, got [%d] instead", n)
	}
	ExistsErr(rt, nil, "error")
	if n := rt.fatalsLen(); n != 1 {
		t.Fatalf("expected 1 fatal, got [%d] instead", n)
	}
}

func TestNoErr(t *testing.T) {
	rt := newRecordingTester()
	NoErr(rt, nil)
	if n := rt.fatalsLen(); n != 0 {
		t.Fatalf("expected 0 fatals, got [%d] instead", n)
	}
	NoErr(rt, errors.New("err"))
	if n := rt.fatalsLen(); n != 1 {
		t.Fatalf("expected 1 fatal, got [%d] instead", n)
	}
}

func TestEqual(t *testing.T) {
	rt := newRecordingTester()
	for i, test := range nonNilTestCases {
		Equal(rt, test.val, test.val, test.name)
		if n := rt.fatalsLen(); n != 0 {
			t.Fatalf("expected 0 fatals, got [%d] instead (iter [%d])", n, i)
		}
	}
}

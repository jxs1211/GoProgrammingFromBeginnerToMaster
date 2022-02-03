package employee

import (
	"errors"
	"testing"
)

type Result struct {
	Count int
}

func (r Result) Int() int { return r.Count }

type Rows []struct{}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...string) (Result, error)
	Query(args []string) (Rows, error)
}

func MaleCount(s Stmt) (int, error) {
	r, err := s.Exec("select * from user where gender=?", "0")
	if err != nil {
		return 0, err
	}
	return r.Int(), nil
}

type FakeMaleStruct struct {
	Stmt
}

func (f FakeMaleStruct) Exec(stmt string, args ...string) (Result, error) {
	return Result{Count: 5}, nil
}

type FakeNilStruct struct {
	Stmt
}

func (f FakeNilStruct) Exec(stmt string, args ...string) (Result, error) {
	return Result{}, errors.New("err is not nil")
}

func TestFakeMaleStruct(t *testing.T) {
	f_male := FakeMaleStruct{}
	i, err := MaleCount(f_male)
	if err != nil {
		t.Errorf("got error when count %#v: %v", f_male, err)
		t.Logf("got error when count %#v: %v", f_male, err)
		return
	}
	want := 5
	if i != want {
		t.Errorf("want: %d, got: %d", want, i)
		t.Logf("want: %d, got: %d", want, i)
		return
	}
}

func TestFakeNilStruct(t *testing.T) {
	f_nil := FakeNilStruct{}
	r, err := MaleCount(f_nil)
	want := errors.New("err is not nil")
	if err == nil {
		t.Errorf("want %s, got: %s", want, err)
		t.Logf("want %s, got: %s", want, err)
	}
	want2 := 0
	if r != 0 {
		t.Errorf("want %d, got: %d", want2, r)
		t.Logf("want %d, got: %d", want2, r)
	}
}

// // 返回女性员工总数
// func FemaleCount(s Stmt) (int, error) {
// 	result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.Int(), nil
// }

// type fakeStmtForMaleCount struct {
// 	Stmt
// }

// func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {
// 	return Result{Count: 5}, nil
// }

// func TestEmployeeMaleCount(t *testing.T) {
// 	f := fakeStmtForMaleCount{}
// 	c, _ := MaleCount(f)
// 	if c != 5 {
// 		t.Errorf("want: %d, actual: %d", 5, c)
// 		return
// 	}
// }

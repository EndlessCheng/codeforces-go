package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func init() { rand.Seed(time.Now().UnixNano()) }

const debugCaseNum = 0
const failedCountLimit = 10

var rg *testutil.RG
var failedCount int

type mockIO struct {
	initData
	answer
	innerData []int

	_t         *testing.T
	caseNum    int
	queryLimit int
	queryCnt   int
}

func (io *mockIO) String() (s string) {
	s = Sprintf("%v", io.innerData)
	//s = strings.Join(io.innerData, "\n")
	return
}

// Mock initData
func (io *mockIO) readInitData() (d initData) {
	return io.initData
}

// Mock query
func (io *mockIO) query(req request) (resp response) {
	if io.caseNum == debugCaseNum {
		Print("Query ", req, " ")
		defer func() { Println(resp) }()
	}

	if io.queryCnt++; io.queryCnt > io.queryLimit {
		io._t.Fatalf("Query Limit Exceeded %d\nInner Data:\n%v", io.caseNum, io)
	}



	return
}

// Check answer
func (io *mockIO) printAnswer(actualAns answer) {
	expectedAns := io.answer
	if !assert.EqualValues(io._t, expectedAns, actualAns, "Wrong Answer %d\nInner Data:\n%v", io.caseNum, io) {
		if failedCount++; failedCount > failedCountLimit {
			io._t.Fatal("too many wrong cases, terminated")
		}
	}

	// for special judge
	ansChecker := func() bool {

		return true
	}
	if !assert.Truef(io._t, ansChecker(), "Wrong Answer %d\nMy Answer:\n%v\nInner Data:\n%v", io.caseNum, actualAns, io) {
		if failedCount++; failedCount > failedCountLimit {
			io._t.Fatal("too many wrong cases, terminated")
		}
	}
}

func Test_doInteraction(_t *testing.T) {
	for tc, checkTC := 1, 1; ; tc++ {
		if tc == debugCaseNum {
			print()
			//debug = true
		}

		io := &mockIO{_t: _t, caseNum: tc}

		rg = testutil.NewRandGenerator()
		n := rg.Int(2, 4)
		a := rg.IntSlice(n, 1, 4)

		io.n = n
		io.ans = a
		io.innerData = a

		io.queryLimit = n * 4

		doInteraction(io)

		if tc == checkTC {
			_t.Logf("%d cases checked.", tc)
			checkTC <<= 1
		}
	}
}

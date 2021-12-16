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
	hiddenData []int //

	_t         *testing.T
	caseNum    int
	queryLimit int
	queryCnt   int
}

func (io *mockIO) String() string {
	hStr := Sprintf("%v", io.hiddenData)
	//hStr := strings.Join(io.hiddenData, "\n")
	return Sprintf("%+v\n%s", io.initData, hStr)
}

// Mock initData
func (io *mockIO) readInitData() (d initData) {
	return io.initData
}

// Check answer
func (io *mockIO) printAnswer(actualAns answer) {
	expectedAns := io.answer
	if !assert.EqualValues(io._t, expectedAns, actualAns, "Wrong Answer %d\nCase Data:\n%v", io.caseNum, io) {
		if failedCount++; failedCount > failedCountLimit {
			io._t.Fatal("too many failed cases, terminated")
		}
	}

	// for special judge
	ansChecker := func() bool {

		return true
	}
	if !assert.Truef(io._t, ansChecker(), "Wrong Answer %d\nMy Answer:\n%v\nCase Data:\n%v", io.caseNum, actualAns, io) {
		if failedCount++; failedCount > failedCountLimit {
			io._t.Fatal("too many failed cases, terminated")
		}
	}
}

// Mock query
func (io *mockIO) query(q request) (resp response) {
	if io.caseNum == debugCaseNum {
		Print("Query ", q, " ")
		defer func() { Println(resp) }()
	}

	io.queryCnt++
	//if io.queryCnt > io.queryLimit { io._t.Fatalf("Query Limit Exceeded %d\nCase Data:\n%v", io.caseNum, io) }

	// calc q ...

	return
}

func Test_doInteraction(_t *testing.T) {
	for tc, checkTC := 1, 1; ; tc++ {
		if tc == debugCaseNum {
			print()
			//debug = true
		}
		io := &mockIO{_t: _t, caseNum: tc}

		// gen random data ...
		rg = testutil.NewRandGenerator()
		n := rg.Int(2, 4)
		a := rg.IntSlice(n, 1, n)

		io.n = n
		io.ans = a
		io.hiddenData = a

		// set limit ...
		io.queryLimit = n + 30

		doInteraction(io)

		if io.queryCnt > io.queryLimit {
			io._t.Errorf("Query Limit Exceeded %d\n%d > %d\nCase Data:\n%v", io.caseNum, io.queryCnt, io.queryLimit, io)
			if failedCount++; failedCount > failedCountLimit {
				io._t.Fatal("too many failed cases, terminated")
			}
		}

		if tc == checkTC {
			_t.Logf("%d cases checked.", tc)
			checkTC <<= 1
		}
	}
}

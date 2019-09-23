package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1133E(t *testing.T) {
	// just copy from website
	rawText := `5 2
1 2 15 15 15
outputCopy
5
inputCopy
6 1
36 4 1 25 9 16
outputCopy
2
inputCopy
4 4
1 10 100 1000
outputCopy
4`
	testutil.AssertEqual(t, rawText, Sol1133E)
}

func TestName(t *testing.T) {
	arr := make([][6]int, 10)
	t.Log(arr)
}
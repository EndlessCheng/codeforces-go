package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol522B(t *testing.T) {
	// just copy from website
	rawText := `
3
1 10
5 5
10 1
outputCopy
75 110 60 
inputCopy
3
2 1
1 2
2 1
outputCopy
6 4 6 `
	testutil.AssertEqualCase(t, rawText, -1, Sol522B)
}

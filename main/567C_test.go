package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol567C(t *testing.T) {
	// just copy from website
	rawText := `
5 2
1 1 2 2 4
outputCopy
4
inputCopy
3 1
1 1 1
outputCopy
1
inputCopy
10 3
1 2 6 2 3 6 9 18 3 9
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, -1, Sol567C)
}

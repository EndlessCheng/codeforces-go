package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol372A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
2
5
7
6
9
8
4
2
outputCopy
5
inputCopy
8
9
1
6
2
6
5
8
3
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, Sol372A)
}

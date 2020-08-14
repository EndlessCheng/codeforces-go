package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1152C(t *testing.T) {
	// just copy from website
	rawText := `
112342324 524224233
outputCopy
299539585
inputCopy
1924 5834
outputCopy
31
inputCopy
6 10
outputCopy
2
inputCopy
21 31
outputCopy
9
inputCopy
5 10
outputCopy
0`
	testutil.AssertEqual(t, rawText, Sol1152C)
}

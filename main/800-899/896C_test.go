package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol896C(t *testing.T) {
	// just copy from website
	rawText := `
10 10 7 9
outputCopy
2
1
0
3
inputCopy
10 10 9 9
outputCopy
1
1
3
3`
	testutil.AssertEqualCase(t, rawText, 0, Sol896C)
}
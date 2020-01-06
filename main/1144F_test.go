package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1144F(t *testing.T) {
	// just copy from website
	rawText := `
6 5
1 5
2 1
1 4
3 1
6 1
outputCopy
YES
10100`
	testutil.AssertEqualCase(t, rawText, 0, Sol1144F)
}

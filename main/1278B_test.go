package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1278B(t *testing.T) {
	// just copy from website
	rawText := `
3
1 3
11 11
30 20
outputCopy
3
0
4`
	testutil.AssertEqualCase(t, rawText, 0, Sol1278B)
}

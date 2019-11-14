package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1257D(t *testing.T) {
	// just copy from website
	rawText := `
2
6
2 3 11 14 1 8
2
3 2
100 1
5
3 5 100 2 3
2
30 5
90 1
outputCopy
5
-1`
	testutil.AssertEqualCase(t, rawText, 0, Sol1257D)
}

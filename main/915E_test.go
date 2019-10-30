package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol915E(t *testing.T) {
	// just copy from website
	rawText := `
4
6
1 2 1
3 4 1
2 3 2
1 3 2
2 4 1
1 4 2
outputCopy
2
0
2
3
1
4`
	testutil.AssertEqualCase(t, rawText, 0, Sol915E)
}

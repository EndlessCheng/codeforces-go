package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol813E(t *testing.T) {
	// just copy from website
	rawText := `
6 2
1 1 1 2 2 2
5
1 6
4 3
1 1
2 6
2 6
outputCopy
2
4
1
3
2`
	testutil.AssertEqualCase(t, rawText, 0, Sol813E)
}

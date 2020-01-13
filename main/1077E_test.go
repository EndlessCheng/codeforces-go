package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1077E(t *testing.T) {
	// just copy from website
	rawText := `
18
2 1 2 10 2 10 10 2 2 1 10 10 10 10 1 1 10 10
outputCopy
14
inputCopy
10
6 6 6 3 6 1000000000 3 3 6 6
outputCopy
9
inputCopy
3
1337 1337 1337
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, Sol1077E)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1265E(t *testing.T) {
	// just copy from website
	rawText := `
1
50
outputCopy
2
inputCopy
3
10 20 50
outputCopy
112`
	testutil.AssertEqualCase(t, rawText, 0, Sol1265E)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol832D(t *testing.T) {
	// just copy from website
	rawText := `
3 2
1 1
1 2 3
2 3 3
outputCopy
2
3
inputCopy
4 1
1 2 3
1 2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, Sol832D)
}

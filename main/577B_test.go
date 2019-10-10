package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol577B(t *testing.T) {
	// just copy from website
	rawText := `
3 5
1 2 3
outputCopy
YES
inputCopy
1 6
5
outputCopy
NO
inputCopy
4 6
3 1 1 3
outputCopy
YES
inputCopy
6 6
5 5 5 5 5 5
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, -1, Sol577B)
}

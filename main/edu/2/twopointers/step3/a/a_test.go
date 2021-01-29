package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 10
1 2 3 4 5 4 3 2 1
outputCopy
3 3
inputCopy
5 6
3 1 1 1 4
outputCopy
5 2
inputCopy
3 100
10 10 10
outputCopy
1 10`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

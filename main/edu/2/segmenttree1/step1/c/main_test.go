package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
3 4 3 5 2
2 0 3
1 1 2
2 0 3
1 0 2
2 0 5
outputCopy
3 2
2 1
2 3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

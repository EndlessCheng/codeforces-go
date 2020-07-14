package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 6
1 2 3 6 5 4 19
1 1 3
1 2 5
1 2 4
2 2 8
1 1 6
1 1 3
outputCopy
3
4
3
6
3`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

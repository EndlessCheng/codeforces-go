package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1366C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2
1 1
0 1
2 3
1 1 0
1 0 0
3 7
1 0 1 1 1 1 1
0 0 0 0 0 0 0
1 1 1 1 1 0 1
3 5
1 0 1 0 0
1 1 1 1 0
0 0 1 0 0
outputCopy
0
3
4
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1366C)
}

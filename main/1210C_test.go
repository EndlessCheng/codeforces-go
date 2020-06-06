package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1210C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 5 6 0 8
1 2
1 3
1 4
4 5
outputCopy
42
inputCopy
7
0 2 3 0 0 0 0
1 2
1 3
2 4
2 5
3 6
3 7
outputCopy
30`
	testutil.AssertEqualCase(t, rawText, 0, CF1210C)
}

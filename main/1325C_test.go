package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
1 3
outputCopy
0
1
inputCopy
6
1 2
1 3
2 4
2 5
5 6
outputCopy
0
3
2
4
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1325C)
}

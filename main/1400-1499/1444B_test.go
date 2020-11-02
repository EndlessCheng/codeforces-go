package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1444B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
1 4
outputCopy
6
inputCopy
2
2 1 2 1
outputCopy
12
inputCopy
3
2 2 2 2 2 2
outputCopy
0
inputCopy
5
13 8 35 94 9284 34 54 69 123 846
outputCopy
2588544`
	testutil.AssertEqualCase(t, rawText, 0, CF1444B)
}

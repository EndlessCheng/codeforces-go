package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1085D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 2
1 3
1 4
outputCopy
2.000000000000000000
inputCopy
6 1
2 1
2 3
2 5
5 4
5 6
outputCopy
0.500000000000000000
inputCopy
5 5
1 2
2 3
3 4
3 5
outputCopy
3.333333333333333333`
	testutil.AssertEqualCase(t, rawText, 0, CF1085D)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1006F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 11
2 1 5
7 10 0
12 6 4
outputCopy
3
inputCopy
3 4 2
1 3 3 3
0 3 3 2
3 0 1 1
outputCopy
5
inputCopy
3 4 1000000000000000000
1 3 3 3
0 3 3 2
3 0 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1006F)
}

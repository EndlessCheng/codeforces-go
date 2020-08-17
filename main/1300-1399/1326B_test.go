package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1326B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 1 1 -2 1
outputCopy
0 1 2 0 3 
inputCopy
3
1000 999999000 -1000000000
outputCopy
1000 1000000000 0 
inputCopy
5
2 1 2 2 3
outputCopy
2 3 5 7 10 `
	testutil.AssertEqualCase(t, rawText, 0, CF1326B)
}

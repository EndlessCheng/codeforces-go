package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1151B(t *testing.T) {
	// just copy from website
	rawText := `
3 2
0 0
0 0
0 0
outputCopy
NIE
inputCopy
2 3
7 7 7
7 7 10
outputCopy
TAK
1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1151B)
}

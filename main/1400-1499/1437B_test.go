package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1437B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
10
4
0110
8
11101000
outputCopy
0
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1437B)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1165D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
8
8 2 12 6 4 24 16 3
1
2
2
2 3
2
4 2
1
3
1
999983
outputCopy
48
4
6
8
9
999966000289`
	testutil.AssertEqualCase(t, rawText, 0, CF1165D)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1278E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2
1 3
3 4
3 5
2 6
outputCopy
9 12
7 10
3 11
1 5
2 4
6 8
inputCopy
1
outputCopy
1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1278E)
}

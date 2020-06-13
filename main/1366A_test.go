package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1366A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 4
1000000000 0
7 15
8 7
outputCopy
2
0
7
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1366A)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1110E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7 2 4 12
7 15 10 12
outputCopy
Yes
inputCopy
3
4 4 4
1 2 3
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1110E)
}

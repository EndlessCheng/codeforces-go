package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF543A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 3 100
1 1 1
outputCopy
10
inputCopy
3 6 5 1000000007
1 2 3
outputCopy
0
inputCopy
3 5 6 11
1 2 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF543A)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1166E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 5
3 1 2 3
3 3 4 5
outputCopy
possible
inputCopy
10 10
1 1
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
outputCopy
impossible`
	testutil.AssertEqualCase(t, rawText, 0, CF1166E)
}

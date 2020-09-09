package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF766E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
1 2
2 3
outputCopy
10
inputCopy
5
1 2 3 4 5
1 2
2 3
3 4
3 5
outputCopy
52
inputCopy
5
10 9 8 7 6
1 2
2 3
3 4
3 5
outputCopy
131`
	testutil.AssertEqualCase(t, rawText, 0, CF766E)
}

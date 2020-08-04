package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF593B(t *testing.T) {
	// just copy from website
	rawText := `
4
1 2
1 2
1 0
0 1
0 2
outputCopy
NO
inputCopy
2
1 3
1 0
-1 3
outputCopy
YES
inputCopy
2
1 3
1 0
0 2
outputCopy
YES
inputCopy
2
1 3
1 0
0 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF593B)
}

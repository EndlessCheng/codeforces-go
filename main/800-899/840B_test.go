package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF840B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 0
1
outputCopy
-1
inputCopy
4 5
0 0 0 -1
1 2
2 3
3 4
1 4
2 4
outputCopy
0
inputCopy
2 1
1 1
1 2
outputCopy
1
1
inputCopy
3 3
0 -1 1
1 2
2 3
1 3
outputCopy
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF840B)
}

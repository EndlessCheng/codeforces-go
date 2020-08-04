package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF711C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 2
0 0 0
1 2
3 4
5 6
outputCopy
10
inputCopy
3 2 2
2 1 2
1 3
2 4
3 5
outputCopy
-1
inputCopy
3 2 2
2 0 0
1 3
2 4
3 5
outputCopy
5
inputCopy
3 2 3
2 1 2
1 3
2 4
3 5
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF711C)
}

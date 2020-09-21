package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF269C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
3 2 10
1 2 10
3 1 5
outputCopy
1
0
1
inputCopy
4 5
1 2 10
1 3 10
2 3 5
4 2 15
3 4 5
outputCopy
0
0
1
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF269C)
}

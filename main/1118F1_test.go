package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1118F1(t *testing.T) {
	// just copy from website
	rawText := `
5
2 0 0 1 2
1 2
2 3
2 4
2 5
outputCopy
1
inputCopy
5
1 0 0 0 2
1 2
2 3
3 4
4 5
outputCopy
4
inputCopy
3
1 1 2
2 3
1 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1118F1)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1343D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 2
1 2 1 2
4 3
1 2 2 1
8 7
6 1 1 7 6 3 4 6
6 6
5 2 6 1 3 4
outputCopy
0
1
4
2
inputCopy
1
4 2
1 2 2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 2, CF1343D)
}

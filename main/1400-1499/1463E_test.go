package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1463E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
2 3 0 5 3
1 5
5 4
outputCopy
3 2 1 5 4
inputCopy
5 2
2 3 0 5 3
1 5
5 1
outputCopy
0
inputCopy
5 1
2 3 0 5 3
4 5
outputCopy
0
inputCopy
5 4
2 3 0 5 3
2 1
3 5
5 2
1 4
outputCopy
3 5 2 1 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1463E)
}

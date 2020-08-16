package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1292B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 2 3 1 0
2 4 20
outputCopy
3
inputCopy
1 1 2 3 1 0
15 27 26
outputCopy
2
inputCopy
1 1 2 3 1 0
2 2 1
outputCopy
0
inputCopy
999 999 100 100 0 0
1000000000000000 1000000000000000 10000000000000000
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1292B)
}

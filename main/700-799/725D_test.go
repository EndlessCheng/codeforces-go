package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF725D(t *testing.T) {
	// just copy from website
	rawText := `
8
20 1000
32 37
40 1000
45 50
16 16
16 16
14 1000
2 1000
outputCopy
3
inputCopy
7
4 4
4 4
4 4
4 4
4 4
4 4
5 5
outputCopy
2
inputCopy
7
14000000003 1000000000000000000
81000000000 88000000000
5000000000 7000000000
15000000000 39000000000
46000000000 51000000000
0 1000000000
0 0
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF725D)
}

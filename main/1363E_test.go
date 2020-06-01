package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1363E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 0 1
20 1 0
300 0 1
4000 0 0
50000 1 0
1 2
2 3
2 4
1 5
outputCopy
4
inputCopy
5
10000 0 1
2000 1 0
300 0 1
40 0 0
1 1 0
1 2
2 3
2 4
1 5
outputCopy
24000
inputCopy
2
109 0 1
205 0 1
1 2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1363E)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5 7
5 7
outputCopy
SMALL
0

inputCopy
2
1 1
300000 1
outputCopy
BIG
299999
inputCopy
2
10 1
13 14
outputCopy
SMALL
6
RPPPRP
inputCopy
3
1 2 1
2 1 2
outputCopy
IMPOSSIBLE
inputCopy
2
2 3
230254780039 337634453113
outputCopy
IMPOSSIBLE
inputCopy
3
1 2 1
994501956360 1410320 1
outputCopy
IMPOSSIBLE`
	testutil.AssertEqualCase(t, rawText, 0, CF1358F)
}

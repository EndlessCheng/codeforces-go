package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1350E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 3
000
111
000
1 1 1
2 2 2
3 3 3
outputCopy
1
1
1
inputCopy
5 2 2
01
10
01
10
01
1 1 4
5 1 4
outputCopy
0
0
inputCopy
5 5 3
01011
10110
01101
11010
10101
1 1 4
1 2 3
5 5 3
outputCopy
1
0
1
inputCopy
1 1 3
0
1 1 1
1 1 2
1 1 3
outputCopy
0
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1350E)
}

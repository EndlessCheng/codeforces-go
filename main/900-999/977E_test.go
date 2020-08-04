package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977E(t *testing.T) {
	// just copy from website
	rawText := `
5 4
1 2
3 4
5 4
3 5
outputCopy
1
inputCopy
17 15
1 8
1 12
5 11
11 9
9 15
15 5
4 13
3 13
4 3
10 16
7 10
16 7
14 3
14 4
17 6
outputCopy
2
inputCopy
4 4
1 2
2 3
1 3
1 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF977E)
}

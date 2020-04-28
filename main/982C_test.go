package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF982C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 4
4 1
3 1
outputCopy
1
inputCopy
3
1 2
1 3
outputCopy
-1
inputCopy
10
7 1
8 4
8 10
4 7
6 5
9 3
3 5
2 10
2 5
outputCopy
4
inputCopy
2
1 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF982C)
}

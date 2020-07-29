package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF358D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 3 4
4 3 2 1
0 1 1 0
outputCopy
13
inputCopy
7
8 5 7 6 1 8 9
2 7 9 5 4 3 1
2 3 3 4 1 1 3
outputCopy
44
inputCopy
3
1 1 1
1 2 1
1 1 1
outputCopy
4
inputCopy
7
1 3 8 9 3 4 4
6 0 6 6 1 8 4
9 6 3 7 8 8 2
outputCopy
42`
	testutil.AssertEqualCase(t, rawText, 0, CF358D)
}

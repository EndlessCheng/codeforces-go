package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF486D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 4
2 1 3 2
1 2
1 3
3 4
outputCopy
8
inputCopy
0 3
1 2 3
1 2
2 3
outputCopy
3
inputCopy
4 8
7 8 7 5 4 6 4 10
1 6
1 2
5 8
1 3
3 5
6 7
3 4
outputCopy
41`
	testutil.AssertEqualCase(t, rawText, 0, CF486D)
}

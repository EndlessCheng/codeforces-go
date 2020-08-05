package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF25D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
0
inputCopy
7
1 2
2 3
3 1
4 5
5 6
6 7
outputCopy
1
3 1 3 7`
	testutil.AssertEqualCase(t, rawText, 0, CF25D)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF940E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5
1 2 3
outputCopy
6
inputCopy
12 10
1 1 10 10 10 10 10 10 9 10 10 10
outputCopy
92
inputCopy
7 2
2 3 6 4 5 7 1
outputCopy
17
inputCopy
8 4
1 3 4 5 5 3 4 1
outputCopy
23`
	testutil.AssertEqualCase(t, rawText, 0, CF940E)
}

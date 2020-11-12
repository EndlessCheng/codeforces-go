package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF609E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7
1 2 3
1 3 1
1 4 5
2 3 2
2 5 3
3 4 2
4 5 4
outputCopy
9
8
11
8
8
8
9`
	testutil.AssertEqualCase(t, rawText, 0, CF609E)
}

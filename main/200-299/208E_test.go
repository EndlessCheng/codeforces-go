package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF208E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
0 1 1 0 4 4
7
1 1
1 2
2 1
2 2
4 1
5 1
6 1
outputCopy
0 0 1 0 0 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF208E)
}

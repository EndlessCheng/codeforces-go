package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1326C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
2 1 3
outputCopy
5 2
inputCopy
5 5
2 1 5 3 4
outputCopy
15 1
inputCopy
7 3
2 7 3 1 5 4 6
outputCopy
18 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1326C)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF528B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 3
3 1
6 1
0 2
outputCopy
3
inputCopy
2
1 5
2 6
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF528B)
}

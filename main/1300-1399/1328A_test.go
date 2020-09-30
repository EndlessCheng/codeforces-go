package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
10 4
13 9
100 13
123 456
92 46
outputCopy
2
5
4
333
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1328A)
}

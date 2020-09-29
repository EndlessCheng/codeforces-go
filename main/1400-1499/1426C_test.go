package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1
5
42
1337
1000000000
outputCopy
0
3
11
72
63244`
	testutil.AssertEqualCase(t, rawText, 0, CF1426C)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1422D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 1 5 5
1 2
4 1
3 3
outputCopy
5
inputCopy
84 5
67 59 41 2
39 56
7 2
15 3
74 18
22 7
outputCopy
42`
	testutil.AssertEqualCase(t, rawText, 0, CF1422D)
}

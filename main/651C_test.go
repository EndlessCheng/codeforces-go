package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF651C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1
7 5
1 5
outputCopy
2
inputCopy
6
0 0
0 1
0 2
-1 1
0 1
1 1
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF651C)
}

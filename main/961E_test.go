package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF961E(t *testing.T) {
	// just copy from website
	rawText := `
5
1 2 3 4 5
outputCopy
0
inputCopy
3
8 12 7
outputCopy
3
inputCopy
3
3 2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF961E)
}

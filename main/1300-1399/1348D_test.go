package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1348D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
9
11
2
outputCopy
3
1 0 2 
3
1 1 2
1
0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1348D)
}

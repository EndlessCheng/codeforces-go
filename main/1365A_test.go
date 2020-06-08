package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1365A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2
0 0
0 0
2 2
0 0
0 1
2 3
1 0 1
1 1 0
3 3
1 0 0
0 0 0
1 0 0
outputCopy
Vivek
Ashish
Vivek
Ashish`
	testutil.AssertEqualCase(t, rawText, 0, CF1365A)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1407A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2
1 0
2
0 0
4
0 1 1 1
4
1 1 0 0
outputCopy
1
0
1
0
2
1 1
4
1 1 0 0`
	testutil.AssertEqualCase(t, rawText, 0, CF1407A)
}

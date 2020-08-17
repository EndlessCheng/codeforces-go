package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1385C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
1 2 3 4
7
4 3 3 8 4 5 2
3
1 1 1
7
1 3 1 4 5 3 2
5
5 4 3 2 3
outputCopy
0
4
0
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1385C)
}

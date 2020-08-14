package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
15 5 3 2
15 5 4 3
15 5 2 1
15 5 5 1
16 7 5 2
20 5 7 3
outputCopy
4
-1
5
2
0
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1183C)
}

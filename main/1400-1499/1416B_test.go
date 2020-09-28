package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1416B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
2 16 4 18
6
1 2 3 4 5 6
5
11 19 1 1 3
outputCopy
2
4 1 2
2 3 3
-1
4
1 2 4
2 4 5
2 3 3
4 5 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1416B)
}

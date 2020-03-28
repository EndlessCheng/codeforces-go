package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8
1 4 8 4 5 6 3 8
16
2 1 3 3 4 3 4 4 1 3 2 2 2 4 1 1
9
2 2 4 4 4 7 7 7 7
outputCopy
3
10
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1183D)
}

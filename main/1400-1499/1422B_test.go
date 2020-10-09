package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1422B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 2
4 2
2 4
4 2
2 4
3 4
1 2 3 4
5 6 7 8
9 10 11 18
outputCopy
8
42`
	testutil.AssertEqualCase(t, rawText, 0, CF1422B)
}

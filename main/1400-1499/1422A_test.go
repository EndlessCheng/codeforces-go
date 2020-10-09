package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1422A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2 3
12 34 56
outputCopy
4
42`
	testutil.AssertEqualCase(t, rawText, 0, CF1422A)
}

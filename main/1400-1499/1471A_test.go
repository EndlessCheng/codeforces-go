package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1471A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3 3
3 6 9
3 3
6 4 11
outputCopy
6 6
7 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1471A)
}

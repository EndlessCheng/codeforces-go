package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1447A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2
3
outputCopy
1
2
5
3 3 3 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1447A)
}

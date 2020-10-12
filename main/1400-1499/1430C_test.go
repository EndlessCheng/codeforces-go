package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1430C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
4
outputCopy
2
2 4
3 3
3 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1430C)
}

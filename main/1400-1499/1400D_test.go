package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1400D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
2 2 2 2 2
6
1 3 3 1 2 3
outputCopy
5
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1400D)
}

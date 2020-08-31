package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1397B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3 2
outputCopy
1
inputCopy
3
1000000000 1000000000 1000000000
outputCopy
1999982505`
	testutil.AssertEqualCase(t, rawText, 0, CF1397B)
}

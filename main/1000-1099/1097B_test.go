package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1097B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
10
20
30
outputCopy
YES
inputCopy
3
10
10
10
outputCopy
NO
inputCopy
3
120
120
120
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1097B)
}

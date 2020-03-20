package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1326A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
2
3
4
outputCopy
-1
57
239
6789`
	testutil.AssertEqualCase(t, rawText, 0, CF1326A)
}

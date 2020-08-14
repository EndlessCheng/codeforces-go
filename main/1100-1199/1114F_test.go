package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1114F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
5 9 1 2
TOTIENT 3 3
TOTIENT 3 4
MULTIPLY 4 4 3
TOTIENT 4 4
outputCopy
1
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1114F)
}

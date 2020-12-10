package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1156C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 3 3 7
outputCopy
2
inputCopy
5 5
10 9 5 8 7
outputCopy
1
inputCopy
3 2
1 4 7
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1156C)
}

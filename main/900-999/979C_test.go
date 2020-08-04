package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF979C(t *testing.T) {
	// just copy from website
	rawText := `
3 1 3
1 2
2 3
outputCopy
5
inputCopy
3 1 3
1 2
1 3
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF979C)
}

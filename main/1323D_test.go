package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1323D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
3
inputCopy
3
1 2 3
outputCopy
2
inputCopy
2
10000000 10000000
outputCopy
20000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1323D)
}

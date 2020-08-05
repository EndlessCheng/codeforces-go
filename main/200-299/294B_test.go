package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF294B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 12
1 3
2 15
2 5
2 1
outputCopy
5
inputCopy
3
1 10
2 1
2 4
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF294B)
}

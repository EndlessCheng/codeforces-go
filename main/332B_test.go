package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF332B(t *testing.T) {
	// just copy from website
	rawText := `
5 2
3 6 1 1 6
outputCopy
1 4
inputCopy
6 2
1 1 1 1 1 1
outputCopy
1 3`
	testutil.AssertEqualCase(t, rawText, 0, CF332B)
}

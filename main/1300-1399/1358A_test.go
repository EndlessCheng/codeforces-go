package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
1 3
2 2
3 3
5 3
outputCopy
1
2
2
5
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1358A)
}

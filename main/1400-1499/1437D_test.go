package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1437D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 4 3 2
2
1 2
3
1 2 3
outputCopy
3
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1437D)
}

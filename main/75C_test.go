package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF75C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 27
3
1 5
10 11
9 11
outputCopy
3
-1
9`
	testutil.AssertEqualCase(t, rawText, 0, CF75C)
}

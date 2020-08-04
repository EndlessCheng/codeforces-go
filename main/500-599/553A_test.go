package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF553A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
2
1
outputCopy
3
inputCopy
4
1
2
3
4
outputCopy
1680`
	testutil.AssertEqualCase(t, rawText, 0, CF553A)
}

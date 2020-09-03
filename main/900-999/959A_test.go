package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF959A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
outputCopy
Ehab
inputCopy
2
outputCopy
Mahmoud`
	testutil.AssertEqualCase(t, rawText, 0, CF959A)
}

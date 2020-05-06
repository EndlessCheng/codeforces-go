package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF482B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1
1 3 3
outputCopy
YES
3 3 3
inputCopy
3 2
1 3 3
1 3 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF482B)
}

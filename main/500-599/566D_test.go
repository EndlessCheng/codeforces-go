package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF566D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 6
3 2 5
1 2 5
3 2 5
2 4 7
2 1 2
3 1 7
outputCopy
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF566D)
}

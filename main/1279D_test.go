package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1279D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 2 1
1 1
outputCopy
124780545
inputCopy
5
2 1 2
2 3 1
3 2 4 3
2 1 4
3 4 3 2
outputCopy
798595483`
	testutil.AssertEqualCase(t, rawText, 0, CF1279D)
}

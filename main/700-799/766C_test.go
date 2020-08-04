package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF766C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
aab
2 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
outputCopy
3
2
2
inputCopy
10
abcdeabcde
5 5 5 5 4 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
outputCopy
401
4
3`
	testutil.AssertEqualCase(t, rawText, 0, CF766C)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1000D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 1 1
outputCopy
2
inputCopy
4
1 1 1 1
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1000D)
}

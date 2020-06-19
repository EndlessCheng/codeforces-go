package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1368A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2 3
5 4 100
outputCopy
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1368A)
}

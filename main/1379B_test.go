package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1379B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 6 13
2 3 1
outputCopy
4 6 5
2 2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1379B)
}

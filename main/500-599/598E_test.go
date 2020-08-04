package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF598E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2 1
2 2 3
2 2 2
2 2 4
outputCopy
5
5
4
0`
	testutil.AssertEqualCase(t, rawText, 0, CF598E)
}

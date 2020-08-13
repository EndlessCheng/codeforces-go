package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF893E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
6 3
4 2
outputCopy
36
6`
	testutil.AssertEqualCase(t, rawText, 0, CF893E)
}
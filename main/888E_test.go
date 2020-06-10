package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF888E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
5 2 4 1
outputCopy
3
inputCopy
3 20
199 41 299
outputCopy
19`
	testutil.AssertEqualCase(t, rawText, 0, CF888E)
}

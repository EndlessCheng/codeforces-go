package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1295A(t *testing.T) {
	// just copy from website
	rawText := `
2
3
4
outputCopy
7
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1295A)
}

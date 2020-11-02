package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1444A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
10 4
12 6
179 822
outputCopy
10
4
179`
	testutil.AssertEqualCase(t, rawText, 0, CF1444A)
}

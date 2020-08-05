package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF652B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 2 1
outputCopy
1 2 1 2
inputCopy
5
1 3 2 2 5
outputCopy
1 5 2 3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF652B)
}

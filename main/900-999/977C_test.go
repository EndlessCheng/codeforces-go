package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 4
3 7 5 1 10 3 20
outputCopy
6
inputCopy
7 2
3 7 5 1 10 3 20
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF977C)
}

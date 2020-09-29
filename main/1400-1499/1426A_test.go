package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7 3
1 5
22 5
987 13
outputCopy
3
1
5
77`
	testutil.AssertEqualCase(t, rawText, 0, CF1426A)
}

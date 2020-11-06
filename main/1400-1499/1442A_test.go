package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1442A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 2 1
5
11 7 9 6 8
5
1 3 1 3 1
4
5 2 1 10
outputCopy
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1442A)
}

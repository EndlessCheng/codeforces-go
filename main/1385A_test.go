package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1385A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2 3
100 100 100
50 49 49
10 30 20
1 1000000000 1000000000
outputCopy
YES
3 2 1
YES
100 100 100
NO
NO
YES
1 1 1000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1385A)
}

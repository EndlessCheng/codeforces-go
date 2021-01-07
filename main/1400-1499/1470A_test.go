package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1470A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5 4
2 3 4 3 2
3 5 12 20
5 5
5 4 3 2 1
10 40 90 160 250
outputCopy
30
190
inputCopy
1
1 1
1
1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1470A)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1311A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 3
10 10
2 4
7 4
9 3
outputCopy
1
0
2
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1311A)
}

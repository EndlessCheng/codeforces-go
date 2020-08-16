package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1299C(t *testing.T) {
	// just copy from website
	rawText := `
4
7 5 5 7
outputCopy
5.666666667
5.666666667
5.666666667
7.000000000
inputCopy
5
7 8 8 10 12
outputCopy
7.000000000
8.000000000
8.000000000
10.000000000
12.000000000
inputCopy
10
3 9 5 5 1 7 5 3 8 7
outputCopy
3.000000000
5.000000000
5.000000000
5.000000000
5.000000000
5.000000000
5.000000000
5.000000000
7.500000000
7.500000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1299C)
}

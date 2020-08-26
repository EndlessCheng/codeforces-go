package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1400B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
33 27
6 10
5 6
100 200
10 10
5 5
1 19
1 3
19 5
outputCopy
11
20
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1400B)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 6
10 15 12
1 9 12 23 26 37
outputCopy
1 1
1 9
2 2
2 13
3 1
3 12
inputCopy
2 3
5 10000000000
5 6 9999999999
outputCopy
1 5
2 1
2 9999999994`
	testutil.AssertEqualCase(t, rawText, 0, CF978C)
}

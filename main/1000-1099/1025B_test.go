package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1025B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
17 18
15 24
12 15
outputCopy
6
inputCopy
2
10 16
7 17
outputCopy
-1
inputCopy
5
90 108
45 105
75 40
165 175
33 30
outputCopy
5
inputCopy
2
6 6
4 9
Parti
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1025B)
}

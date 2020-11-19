package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1440B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 4
0 24 34 58 62 64 69 78
2 2
27 61 81 91
4 3
2 4 16 18 21 27 36 53 82 91 92 95
3 4
3 11 12 22 33 35 38 67 69 71 94 99
2 1
11 41
3 3
1 1 1 1 1 1 1 1 1
outputCopy
165
108
145
234
11
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1440B)
}

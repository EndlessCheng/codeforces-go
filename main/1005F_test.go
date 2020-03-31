package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1005F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 3
1 2
2 3
1 4
4 3
outputCopy
2
1110
1011
inputCopy
4 6 3
1 2
2 3
1 4
4 3
2 4
1 3
outputCopy
1
101001
inputCopy
5 6 2
1 2
1 3
2 4
2 5
3 4
3 5
outputCopy
2
111100
110110`
	testutil.AssertEqualCase(t, rawText, 0, CF1005F)
}

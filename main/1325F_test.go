package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 6
1 3
3 4
4 2
2 6
5 6
5 1
outputCopy
1
1 6 4
inputCopy
6 8
1 3
3 4
4 2
2 6
5 6
5 1
1 4
2 5
outputCopy
2
4
1 5 2 4
inputCopy
5 4
1 2
1 3
2 4
2 5
outputCopy
1
3 4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1325F)
}

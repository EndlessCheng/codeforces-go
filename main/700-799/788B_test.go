package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF788B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
1 2
1 3
1 4
1 5
outputCopy
6
inputCopy
5 3
1 2
2 3
4 5
outputCopy
0
inputCopy
2 2
1 1
1 2
outputCopy
1
inputCopy
4 4
2 3
2 4
3 4
4 4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF788B)
}

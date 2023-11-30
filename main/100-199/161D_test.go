package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF161D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
1 2
2 3
3 4
2 5
outputCopy
4
inputCopy
5 3
1 2
2 3
3 4
4 5
outputCopy
2
inputCopy
10 3
2 1
3 1
4 3
5 4
6 5
7 1
8 6
9 2
10 6
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF161D)
}

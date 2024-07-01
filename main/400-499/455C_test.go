package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF455C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 0 6
2 1 2
2 3 4
2 5 6
2 3 2
2 5 3
1 1
outputCopy
4
inputCopy
10 6 7
1 2
2 3
1 4
3 5
1 6
5 7
1 1
2 8 10
1 10
2 9 5
1 7
1 4
2 1 9
outputCopy
5
1
5
5`
	testutil.AssertEqualCase(t, rawText, 2, cf455C)
}

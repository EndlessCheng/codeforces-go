package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abacaba
6
4 7
1 2
1 1
3 4
5 7
1 3
outputCopy
1 1
1 2
1 3
5 7
3 4
4 7
inputCopy
a
1
1 1
outputCopy
1 1
inputCopy
aa
5
1 1
1 2
2 2
1 1
2 2
outputCopy
1 1
1 1
2 2
2 2
1 2`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

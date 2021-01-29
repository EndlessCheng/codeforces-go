package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
2
1 3
2
3 4
2
2 3
outputCopy
3 3 3 3
inputCopy
1
5
4
3 6 7 10
4
18 3 9 11
1
20
outputCopy
5 6 9 20
inputCopy
1
4
1
3
1
1
1
2
outputCopy
4 3 1 2`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

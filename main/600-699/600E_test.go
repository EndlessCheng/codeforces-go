package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF600E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2 3 4
1 2
2 3
2 4
outputCopy
10 9 3 4
inputCopy
15
1 2 3 1 2 3 3 1 1 3 2 2 1 2 3
1 2
1 3
1 4
1 14
1 15
2 5
2 6
2 7
3 8
3 9
3 10
4 11
4 12
4 13
outputCopy
6 5 4 3 2 3 3 1 1 3 2 2 1 2 3`
	testutil.AssertEqualCase(t, rawText, -1, CF600E)
}

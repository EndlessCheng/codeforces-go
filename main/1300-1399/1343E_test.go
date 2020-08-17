package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1343E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 3 2 3 4
1 2 3
1 2
1 3
1 4
7 9 1 5 7
2 10 4 8 5 6 7 3 3
1 2
1 3
1 4
3 2
3 5
4 2
5 6
1 7
6 7
outputCopy
7
12`
	testutil.AssertEqualCase(t, rawText, 0, CF1343E)
}

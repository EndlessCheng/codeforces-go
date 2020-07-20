package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1385E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 1
0 1 3
5 5
0 2 1
1 1 5
1 5 4
0 5 2
1 3 5
4 5
1 1 2
0 4 3
1 3 1
0 2 3
1 2 4
4 5
1 4 1
1 1 3
0 1 2
1 2 4
1 3 2
outputCopy
YES
3 1
YES
2 1
1 5
5 4
2 5
3 5
YES
1 2
3 4
3 1
3 2
2 4
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1385E)
}

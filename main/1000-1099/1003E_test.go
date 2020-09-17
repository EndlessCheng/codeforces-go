package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1003E(t *testing.T) {
	// just copy from website
	rawText := `
6 3 3
outputCopy
YES
3 1
4 1
1 2
5 2
2 6
inputCopy
6 2 3
outputCopy
NO
inputCopy
10 4 3
outputCopy
YES
2 9
2 10
10 3
3 1
6 10
8 2
4 3
5 6
6 7
inputCopy
8 5 3
outputCopy
YES
2 5
7 2
3 7
3 1
1 6
8 7
4 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1003E)
}

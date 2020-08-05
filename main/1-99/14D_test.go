package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF14D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
2 3
3 4
outputCopy
1
inputCopy
7
1 2
1 3
1 4
1 5
1 6
1 7
outputCopy
0
inputCopy
6
1 2
2 3
2 4
5 4
6 4
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF14D)
}

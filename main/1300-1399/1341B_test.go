package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1341B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
8 6
1 2 4 1 2 4 1 2
5 3
3 2 3 2 1
10 4
4 3 4 3 2 3 2 1 0 1
15 7
3 7 4 8 2 3 4 5 21 2 3 4 2 1 3
7 5
1 2 3 4 5 6 1
outputCopy
3 2
2 2
2 1
3 1
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1341B)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF954D(t *testing.T) {
	// just copy from website
	rawText := `
5 4 1 5
1 2
2 3
3 4
4 5
outputCopy
0
inputCopy
5 4 3 5
1 2
2 3
3 4
4 5
outputCopy
5
inputCopy
5 6 1 5
1 2
1 3
1 4
4 5
3 5
2 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF954D)
}

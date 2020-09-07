package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF796D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2 4
1 6
1 2
2 3
3 4
4 5
5 6
outputCopy
1
5
inputCopy
6 3 2
1 5 6
1 2
1 3
1 4
1 5
5 6
outputCopy
2
4 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF796D)
}

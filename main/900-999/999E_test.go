package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF999E(t *testing.T) {
	// just copy from website
	rawText := `
9 9 1
1 2
1 3
2 3
1 5
5 6
6 1
1 8
9 8
7 1
outputCopy
3
inputCopy
5 4 5
1 2
2 3
3 4
4 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF999E)
}

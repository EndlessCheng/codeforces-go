package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1352B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
10 3
100 4
8 7
97 2
8 8
3 10
5 3
1000000000 9
outputCopy
YES
4 2 4
YES
55 5 5 35
NO
NO
YES
1 1 1 1 1 1 1 1
NO
YES
3 1 1
YES
111111110 111111110 111111110 111111110 111111110 111111110 111111110 111111110 111111120`
	testutil.AssertEqualCase(t, rawText, 0, CF1352B)
}

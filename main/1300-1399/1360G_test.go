package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1360G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 6 2 1
2 2 2 1
2 2 2 2
4 4 2 2
2 1 1 2
outputCopy
YES
010001
100100
001010
NO
YES
11
11
YES
1100
1100
0011
0011
YES
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1360G)
}

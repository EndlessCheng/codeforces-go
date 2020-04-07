package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1101D(t *testing.T) {
	// just copy from website
	rawText := `
input
3
2 3 4
1 2
2 3
output
1
input
3
2 3 4
1 3
2 3
output
2
input
3
1 1 1
1 2
2 3
output
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1101D)
}

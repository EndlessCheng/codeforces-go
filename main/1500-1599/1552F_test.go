package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1552/F
// https://codeforces.com/problemset/status/1552/problem/F
func TestCF1552F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 2 0
6 5 1
7 4 0
8 1 1
outputCopy
23
inputCopy
1
454971987 406874902 1
outputCopy
503069073
inputCopy
5
243385510 42245605 0
644426565 574769163 0
708622105 208990040 0
786625660 616437691 0
899754846 382774619 0
outputCopy
899754847
inputCopy
5
200000000 100000000 1
600000000 400000000 0
800000000 300000000 0
900000000 700000000 1
1000000000 500000000 0
outputCopy
3511295`
	testutil.AssertEqualCase(t, rawText, 0, CF1552F)
}

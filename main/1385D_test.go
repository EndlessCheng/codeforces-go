package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1385D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
8
bbdcaaaa
8
asdfghjk
8
ceaaaabb
8
bbaaddcc
1
z
2
ac
outputCopy
0
7
4
5
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1385D)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1342A(t *testing.T) {
	// just copy from website
	rawText := `
input
2
1 3
391 555
0 0
9 4
output
1337
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1342A)
}

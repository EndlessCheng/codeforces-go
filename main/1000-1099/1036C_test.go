package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1036C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1000
1024 1024
65536 65536
999999 1000001
outputCopy
1000
1
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1036C)
}

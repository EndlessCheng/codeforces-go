package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1301D(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	// just copy from website
	rawText := `
2 3 14
outputCopy
YES
9
1 D
1 RLU
1 R
1 D
1 RLU
1 R
1 D
1 U
2 L
inputCopy
3 3 4
outputCopy
YES
2
2 R
2 L
inputCopy
3 3 1000000000
outputCopy
NO
inputCopy
3 3 8
outputCopy
YES
3
2 R
2 D
1 LLRR
inputCopy
4 4 9
outputCopy
YES
1
3 RLD
inputCopy
3 4 16
outputCopy
YES
8
3 R
3 L
1 D
3 R
1 D
1 U
3 L
1 D`
	testutil.AssertEqualCase(t, rawText, 1, CF1301D)
}

package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1304E(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	// just copy from website
	rawText := `
5
1 2
2 3
3 4
4 5
5
1 3 1 2 2
1 4 1 3 2
1 4 1 3 3
4 2 3 3 9
5 2 3 3 9
outputCopy
YES
YES
NO
YES
NO
inputCopy
3
1 2
2 3
36
1 3 2 2 3
1 3 1 1 1
1 3 1 1 2
1 3 1 1 3
1 3 1 1 4
1 3 1 1 5
1 3 1 1 6
1 3 1 1 999999997
1 3 1 1 999999998
1 3 1 1 999999999
1 3 1 1 1000000000
3 1 1 1 1
3 1 1 1 2
3 1 1 1 3
3 1 1 1 4
3 1 1 1 5
3 1 1 1 6
3 1 1 1 999999997
3 1 1 1 999999998
3 1 1 1 999999999
3 1 1 1 1000000000
1 3 1 2 1
1 3 1 2 2
1 3 1 2 3
1 3 1 2 4
1 3 1 2 5
1 3 1 2 6
1 3 1 2 999999997
1 3 1 2 999999998
1 3 1 2 999999999
1 3 1 2 1000000000
3 1 1 2 1
3 1 1 2 2
3 1 1 2 3
3 1 1 2 4
3 1 1 2 5
3 1 1 2 6
outputCopy
YES
NO
YES
YES
YES
YES
YES
YES
YES
YES
YES
NO
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES`
	testutil.AssertEqualCase(t, rawText, 2, CF1304E)
}

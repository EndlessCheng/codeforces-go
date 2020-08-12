package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1074D(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	// just copy from website
	rawText := `
12
2 1 2
2 1 1073741822
1 0 3 4
2 0 0
2 3 3
2 0 3
1 6 7 3
2 4 4
1 0 2 1
2 0 0
2 4 4
2 0 0
outputCopy
-1
-1
-1
-1
5
-1
6
3
5
inputCopy
4
1 5 5 9
1 6 6 5
1 6 5 10
2 6 5
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, 1, CF1074D)
}

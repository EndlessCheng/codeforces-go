package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1354/E
// https://codeforces.com/problemset/status/1354/problem/E
func TestCF1354E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
2 2 2
3 1
5 4
2 5
outputCopy
YES
112323
inputCopy
5 9
0 2 3
1 2
1 3
1 5
2 3
2 4
2 5
3 4
3 5
4 5
outputCopy
NO
inputCopy
4 0
0 4 0
outputCopy
YES
2222
inputCopy
8 6
4 4 0
1 2
2 3
3 4
5 6
5 7
5 8
outputCopy
NO` // 自己构造的一个卡背包的反例
	testutil.AssertEqualCase(t, rawText, 0, CF1354E)
}

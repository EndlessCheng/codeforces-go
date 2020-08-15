package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://ac.nowcoder.com/acm/problem/51178
func Test_nc51178(t *testing.T) {
	customTestCases := [][2]string{
		{
			`7
1
1
1
1
1
1
1
1 3
2 3
6 4
7 4
4 5
3 5
0 0`,
			`5`,
		},
		// TODO: 测试参数的下界和上界

	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, nc51178)
}

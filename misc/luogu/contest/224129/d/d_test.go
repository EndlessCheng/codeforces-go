package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var customTestCases = [][2]string{
	{
		`3 6 2
1 5 1
5 6 2
5 6 3
`,
		`5`,
	},
	{
		`4 7 3
1 7 1
2 5 4
4 7 5
1 2 10`,
		`9`,
	},
}

func Test_d(t *testing.T) {
	if len(customTestCases[0][0]) == 0 {
		t.Error("请添加测试数据！")
		return
	}
	testutil.AssertEqualStringCaseWithPrefix(t, customTestCases, 0, run, "")
}

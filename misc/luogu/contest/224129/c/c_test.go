package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var customTestCases = [][2]string{
	{
		`8 3
vvvaannn
4
3
5`,
		`18
15
12`,
	},
	{
		`3 4
van
1
1
1
1`,
		``,
	},
}

func Test_c(t *testing.T) {
	if len(customTestCases[0][0]) == 0 {
		t.Error("请添加测试数据！")
		return
	}
	testutil.AssertEqualStringCaseWithPrefix(t, customTestCases, 0, run, "")
}

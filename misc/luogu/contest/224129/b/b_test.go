package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var customTestCases = [][2]string{
	{
		`3 3
0 0 0 
1 0 1
0 1 1`,
		`2`,
	},
	{
		`4 3
0 0 0 
1 0 1
0 1 1
1 1 1`,
		`0`,
	},
}

func Test_b(t *testing.T) {
	if len(customTestCases[0][0]) == 0 {
		t.Error("请添加测试数据！")
		return
	}
	testutil.AssertEqualStringCaseWithPrefix(t, customTestCases, 0, run, "")
}

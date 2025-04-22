package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var customTestCases = [][2]string{
	{
		`3 2
1 2
5 7
8 9`,
		`1`,
	},
}

func Test_a(t *testing.T) {
	if len(customTestCases[0][0]) == 0 {
		t.Error("请添加测试数据！")
		return
	}
	testutil.AssertEqualStringCaseWithPrefix(t, customTestCases, 0, run, "")
}

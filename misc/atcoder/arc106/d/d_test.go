// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc106/submit?taskScreenName=arc106_d
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`3 3
1 2 3`,
			`12
50
216`,
		},
		{
			`10 10
1 1 1 1 1 1 1 1 1 1`,
			`90
180
360
720
1440
2880
5760
11520
23040
46080`,
		},
		{
			`2 5
1234 5678`,
			`6912
47775744
805306038
64822328
838460992`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc106/tasks/arc106_d

// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc275/submit?taskScreenName=abc275_f
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`4 5
1 2 3 4`,
			`1
2
1
1
1`,
		},
		{
			`1 5
3`,
			`-1
-1
0
-1
-1`,
		},
		{
			`12 20
2 5 6 5 2 1 7 9 7 2 5 5`,
			`2
1
2
2
1
2
1
2
2
1
2
1
1
1
2
2
1
1
1
1`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc275/tasks/abc275_f

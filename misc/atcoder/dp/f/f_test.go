// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/dp/submit?taskScreenName=dp_f
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`axyb
abyxb`,
			`axb`,
		},
		{
			`aa
xayaz`,
			`aa`,
		},
		{
			`a
z`,
			``,
		},
		{
			`abracadabra
avadakedavra`,
			`aaadara`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/dp/tasks/dp_f

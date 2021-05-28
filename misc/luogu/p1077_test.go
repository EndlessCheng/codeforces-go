package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1077
func Test_p1077(t *testing.T) {
	customTestCases := [][2]string{
		{
			`2 4
3 2`,
			`2`,
		},
	}

	tarCase := 0 // -1
	testutil.AssertEqualStringCase(t, customTestCases, tarCase, p1077)
}

// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P10247
func Test_p10247(t *testing.T) {
	testCases := [][2]string{
		{
			`6 5
1 2
2 3
1 3
2 5
3 6`,
			`5 0 4 5 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p10247)
}

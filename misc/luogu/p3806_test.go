// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P3806
func Test_p3806(t *testing.T) {
	testCases := [][2]string{
		{
			`2 1
1 2 2
2`,
			`AYE`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p3806)
}

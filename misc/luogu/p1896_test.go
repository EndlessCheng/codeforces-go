// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1896
func Test_p1896(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2`,
			`16`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1896)
}

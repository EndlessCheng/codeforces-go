// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P3628
func Test_p3628(t *testing.T) {
	testCases := [][2]string{
		{
			`4 
-1 10 -20 
2 2 3 4`,
			`9`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p3628)
}

// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P4316
func Test_p4316(t *testing.T) {
	testCases := [][2]string{
		{
			`4 4 
1 2 1 
1 3 2 
2 3 3 
3 4 4`,
			`7.00`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p4316)
}

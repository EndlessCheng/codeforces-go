// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P2590
func Test_p2590(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1 2
2 3
4 1
4 2 1 3
12
QMAX 3 4
QMAX 3 3
QMAX 3 2
QMAX 2 3
QSUM 3 4
QSUM 2 1
CHANGE 1 5
QMAX 3 4
CHANGE 3 6
QMAX 3 4
QMAX 2 4
QSUM 3 4`,
			`4
1
2
2
10
6
5
6
5
16`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p2590)
}

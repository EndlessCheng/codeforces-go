// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	t.Log("记得初始化所有全局变量")
	targetCaseNum := -1
	if err := testutil.RunLeetCodeClassWithFile(t, Constructor, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-76/problems/design-an-atm-machine/

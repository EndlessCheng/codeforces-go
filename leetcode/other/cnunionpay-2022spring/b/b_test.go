// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	t.Log("记得初始化所有全局变量")
	examples := [][3]string{
		{
			`["DiscountSystem","addActivity","consume","removeActivity","consume"]`,
			`[[],[1,15,5,7,2],[101,16],[1],[102,19]]`,
			`[null,null,11,null,19]`,
		},
		{
			`["DiscountSystem","addActivity","addActivity","consume","removeActivity","consume","consume","consume","consume"]`,
			`[[],[1,10,6,3,2],[2,15,8,8,2],[101,13],[2],[101,17],[101,11],[102,16],[102,21]]`,
			`[null,null,null,7,null,11,11,10,21]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeClassWithExamples(t, Constructor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/cnunionpay-2022spring/problems/kDPV0f/

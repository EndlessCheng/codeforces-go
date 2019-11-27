package copypasta

import "testing"

func Test_convolution(t *testing.T) {
	// 使用 https://www.luogu.com.cn/problem/P3803 测试
	res := convolution([]int64{1, 2}, []int64{1, 2, 1})
	t.Log(res)
}

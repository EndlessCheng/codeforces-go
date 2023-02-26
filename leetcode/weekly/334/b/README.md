### 前置知识：模运算的性质

设 $a=k_1m+r_1,b=k_2m+r_2$。

那么 $(a+b)\bmod m = (r_1+r_2)\bmod m = (a\bmod m + b\bmod m)\bmod m$。

### 思路

从左到右计算。设当前数字为 $x$，每遇到一个数字 $d$，就把 $x$ 更新为 $(10x+d)\bmod m$。

附：[视频讲解](https://www.bilibili.com/video/BV1wj411G7sH/)

```py [sol1-Python3]
class Solution:
    def divisibilityArray(self, word: str, m: int) -> List[int]:
        ans, x = [], 0
        for d in map(int, word):
            x = (x * 10 + d) % m
            ans.append(int(x == 0))
        return ans
```

```go [sol1-Go]
func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	for i, c := range word {
		x = (x*10 + int(c-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$O(1)$。返回值不计入。

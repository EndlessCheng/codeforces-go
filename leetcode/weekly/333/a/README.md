做法和归并排序是一样的。

附：[视频讲解](https://www.bilibili.com/video/BV1jM411J7y7/)

```py [sol1-Python3]
class Solution:
    def mergeArrays(self, a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
        ans = []
        i, n = 0, len(a)
        j, m = 0, len(b)
        while True:
            if i == n:
                ans.extend(b[j:])
                return ans
            if j == m:
                ans.extend(a[i:])
                return ans
            if a[i][0] < b[j][0]:
                ans.append(a[i])
                i += 1
            elif a[i][0] > b[j][0]:
                ans.append(b[j])
                j += 1
            else:
                a[i][1] += b[j][1]
                ans.append(a[i])
                i += 1
                j += 1
```

```go [sol1-Go]
func mergeArrays(a, b [][]int) (ans [][]int) {
	i, n := 0, len(a)
	j, m := 0, len(b)
	for {
		if i == n {
			return append(ans, b[j:]...)
		}
		if j == m {
			return append(ans, a[i:]...)
		}
		if a[i][0] < b[j][0] {
			ans = append(ans, a[i])
			i++
		} else if a[i][0] > b[j][0] {
			ans = append(ans, b[j])
			j++
		} else {
			a[i][1] += b[j][1]
			ans = append(ans, a[i])
			i++
			j++
		}
	}
}
```

### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$O(1)$。不计入返回值，仅用到若干额外变量。

[本题视频讲解](https://www.bilibili.com/video/BV1pC4y1j7Pw/)

思路类似冒泡排序，毕竟这题就是把 $s$ 排序。

对于每个 $0$，它左边有多少个 $1$，就移动多少次。

所以一边遍历 $s$，一边统计 $1$ 的个数 $\textit{cnt}_1$，遇到 $0$ 就把 $\textit{cnt}_1$ 加入答案。

> 也可以用逆序对来思考。

```py [sol-Python3]
class Solution:
    def minimumSteps(self, s: str) -> int:
        ans = cnt1 = 0
        for c in s:
            if c == '1':
                cnt1 += 1
            else:
                ans += cnt1
        return ans
```

```java [sol-Java]
class Solution {
    public long minimumSteps(String s) {
        long ans = 0;
        int cnt1 = 0;
        for (char c : s.toCharArray()) {
            if (c == '1') {
                cnt1++;
            } else {
                ans += cnt1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumSteps(string s) {
        long long ans = 0;
        int cnt1 = 0;
        for (char c : s) {
            if (c == '1') {
                cnt1++;
            } else {
                ans += cnt1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumSteps(s string) (ans int64) {
	cnt1 := 0
	for _, c := range s {
		if c == '1' {
			cnt1++
		} else {
			ans += int64(cnt1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

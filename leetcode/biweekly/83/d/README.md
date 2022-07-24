下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

考虑包含 $1$ 到 $k$ 的最短前缀，无法得到的子序列的第一个数必然在里面。

#### 提示 2-1

取这个前缀的最后一个数当做子序列的第一个数。

#### 提示 2-2

这个数有个特殊性质，它在这个前缀中只出现一次。

#### 提示 3

去掉这个前缀，考虑下一个包含 $1$ 到 $k$ 的最短前缀。在提示 2-1 的前提下，子序列的第二个数必然在这个前缀中，同样地，取前缀的最后一个数当做子序列的第二个数。

根据提示 2-2，按照这种取法，取到的这两个数组成的子序列，一定不会位于第一个前缀中（读者可以用这两个数相同和不同来分类讨论）。因此这种取法是正确的。

#### 提示 4

不断重复这一过程，直到剩余部分无法包含 $1$ 到 $k$ 时停止。

设我们取到了 $m$ 个数，对应着 $\textit{rolls}$ 的 $m$ 个子段。由于每一段都包含 $1$ 到 $k$，$\textit{rolls}$ 必然包含长度为 $m$ 的子序列：每一段都选一个元素即可组成这样的子序列。

因此答案至少为 $m+1$。

我们可以构造出一个长为 $m+1$ 的子序列，它不在 $\textit{rolls}$ 中：前 $m$ 个数取每个子段的最后一个数，第 $m+1$ 个数取不在剩余部分中的数。

代码实现时，可以用一个 $\textit{mark}$ 数组标记每个元素属于哪个子段。

#### 复杂度分析

- 时间复杂度：$O(n+k)$，其中 $n$ 为 $\textit{rolls}$ 的长度。
- 空间复杂度：$O(k)$。

```py [sol1-Python3]
class Solution:
    def shortestSequence(self, rolls: List[int], k: int) -> int:
        mark = [0] * (k + 1)  # mark[v] 标记 v 属于哪个子段
        ans, left = 1, k
        for v in rolls:
            if mark[v] < ans:
                mark[v] = ans
                left -= 1
                if left == 0:
                    left = k
                    ans += 1
        return ans
```

```java [sol1-Java]
class Solution {
    public int shortestSequence(int[] rolls, int k) {
        var mark = new int[k + 1]; // mark[v] 标记 v 属于哪个子段
        int ans = 1, left = k;
        for (var v : rolls)
            if (mark[v] < ans) {
                mark[v] = ans;
                if (--left == 0) {
                    left = k;
                    ++ans;
                }
            }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int shortestSequence(vector<int> &rolls, int k) {
        int ans = 1, left = k, mark[k + 1]; // mark[v] 标记 v 属于哪个子段
        memset(mark, 0, sizeof(mark));
        for (int v : rolls)
            if (mark[v] < ans) {
                mark[v] = ans;
                if (--left == 0) {
                    left = k;
                    ++ans;
                }
            }
        return ans;
    }
};
```

```go [sol1-Go]
func shortestSequence(rolls []int, k int) int {
	mark := make([]int, k+1) // mark[v] 标记 v 属于哪个子段
	ans, left := 1, k
	for _, v := range rolls {
		if mark[v] < ans {
			mark[v] = ans
			if left--; left == 0 {
				left = k
				ans++
			}
		}
	}
	return ans
}
```

#### 思考题


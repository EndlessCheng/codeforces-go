注意到，如果子串中的 $0$ 非常多，多到 $0$ 的个数的平方比 $1$ 的个数都要大，那么这样的子串必然不是 $1$ 显著子串。

设 $\textit{cnt}_0$ 为子串中的 $0$ 的个数，$\textit{cnt}_1$ 为子串中的 $1$ 的个数，那么必须满足

$$
\textit{cnt}_0^2 \le \textit{cnt}_1 \le n
$$

解得

$$
\textit{cnt}_0 \le \sqrt{n}
$$

所以子串中的 $0$ 的个数不会超过 $\sqrt{n}$。

据此：

- 枚举子串左端点 $\textit{left}$。
- 枚举子串中有**恰好** $\textit{cnt}_0 = 0,1,2,3,\cdots$ 个 $0$，如果 $\textit{cnt}_0^2$ 超过了 $s$ 中的 $1$ 的个数，则退出循环。这只会枚举 $\mathcal{O}(\sqrt{n})$ 次。

定义：

- 子串左端点为 $\textit{left}$。
- 子串中**恰好**有 $\textit{cnt}_0$ 个 $0$。那么子串右端点的下标范围就是 $[p,q-1]$，其中 $p$ 为从 $\textit{left}$ 开始的第 $\textit{cnt}_0$ 个 $0$ 的下标，$q$ 为从 $\textit{left}$ 开始的第 $\textit{cnt}_0+1$ 个 $0$ 的下标。
- 从子串左端点 $\textit{left}$ 到 $p$ 之间有 $\textit{cnt}_1$ 个 $1$。

分类讨论：

- 如果 $\textit{cnt}_0^2\le \textit{cnt}_1$，那么子串右端点可以是 $p,p+1,p+2,\cdots, q-1$，一共有 $q-p$ 个。注意一定要保证子串中**恰好**有 $\textit{cnt}_0$ 个 $0$。
- 如果 $\textit{cnt}_0^2> \textit{cnt}_1$，那么为了补足 $1$ 的个数，子串右端点的最小值就不是 $p$ 了，而是 $p + (\textit{cnt}_0^2 - \textit{cnt}_1)$，一共有 $q - p - (\textit{cnt}_0^2 - \textit{cnt}_1)$ 个。如果这个值是负数，则说明没有符合要求的子串。

综上所述，当子串左端点为 $\textit{left}$ 且子串中有恰好 $\textit{cnt}_0$ 个 $0$ 时，一共有

$$
\max(q-p-d, 0)
$$

个 $1$ 显著子串。其中 

$$
d = \max(\textit{cnt}_0^2 - \textit{cnt}_1, 0)
$$

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Mi421a7cZ/) 第三题，欢迎点赞关注！

## 写法一：枚举左端点

```py [sol-Python3]
class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        n = len(s)
        a = [i for i, b in enumerate(s) if b == '0']
        tot1 = n - len(a)
        a.append(n)  # 哨兵
        ans = i = 0  # >= left 的第一个 0 的下标是 a[i]
        for left, b in enumerate(s):
            if b == '1':
                ans += a[i] - left  # 不含 0 的子串个数
            for k in range(i, len(a) - 1):
                cnt0 = k - i + 1
                if cnt0 * cnt0 > tot1:
                    break
                cnt1 = a[k] - left - (k - i)
                # 可以改成手动比大小，那样更快
                ans += max(a[k + 1] - a[k] - max(cnt0 * cnt0 - cnt1, 0), 0)
            if b == '0':
                i += 1  # 这个 0 后面不会再枚举到了
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSubstrings(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int m = 0;
        int[] a = new int[n + 1];
        for (int i = 0; i < n; i++) {
            if (s[i] == '0') {
                a[m++] = i;
            }
        }

        int tot1 = n - m;
        a[m] = n; // 哨兵

        int ans = 0;
        int i = 0; // >= left 的第一个 0 的下标是 a[i]
        for (int left = 0; left < n; left++) {
            if (s[left] == '1') {
                ans += a[i] - left; // 不含 0 的子串个数
            }
            for (int k = i; k < m; k++) {
                int cnt0 = k - i + 1;
                if (cnt0 * cnt0 > tot1) {
                    break;
                }
                int cnt1 = a[k] - left - (k - i);
                ans += Math.max(a[k + 1] - a[k] - Math.max(cnt0 * cnt0 - cnt1, 0), 0);
            }
            if (s[left] == '0') {
                i++; // 这个 0 后面不会再枚举到了
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSubstrings(string s) {
        int n = s.length();
        vector<int> a;
        for (int i = 0; i < n; i++) {
            if (s[i] == '0') {
                a.push_back(i);
            }
        }

        int tot1 = n - a.size();
        a.push_back(n); // 哨兵

        int ans = 0, i = 0; // >= left 的第一个 0 的下标是 a[i]
        for (int left = 0; left < n; left++) {
            if (s[left] == '1') {
                ans += a[i] - left; // 不含 0 的子串个数
            }
            for (int k = i; k < a.size() - 1; k++) {
                int cnt0 = k - i + 1;
                if (cnt0 * cnt0 > tot1) {
                    break;
                }
                int cnt1 = a[k] - left - (k - i);
                ans += max(a[k + 1] - a[k] - max(cnt0 * cnt0 - cnt1, 0), 0);
            }
            if (s[left] == '0') {
                i++; // 这个 0 后面不会再枚举到了
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubstrings(s string) (ans int) {
	a := []int{}
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}

	n := len(s)
	tot1 := n - len(a)
	a = append(a, n) // 哨兵

	for left, b := range s {
		if b == '1' {
			ans += a[0] - left // 不含 0 的子串个数
		}
		for k, j := range a[:len(a)-1] {
			cnt0 := k + 1
			if cnt0*cnt0 > tot1 {
				break
			}
			cnt1 := j - left - k
			ans += max(a[k+1]-j-max(cnt0*cnt0-cnt1, 0), 0)
		}
		if b == '0' {
			a = a[1:] // 这个 0 后面不会再枚举到了
		}
	}
	return
}
```

## 写法二：枚举右端点

改成枚举子串右端点 $\textit{right}$，做法同上，但可以在枚举的同时，维护 $0$ 的下标列表 $a$。

```py [sol-Python3]
class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        ans = tot1 = 0
        a = [-1]  # 哨兵
        for right, b in enumerate(s):
            if b == '0':
                a.append(right)
            else:
                ans += right - a[-1]  # 不含 0 子串的个数
                tot1 += 1
            for k in range(len(a) - 1, 0, -1):
                cnt0 = len(a) - k
                if cnt0 * cnt0 > tot1:
                    break
                cnt1 = right - a[k] + 1 - cnt0
                # 可以改成手动比大小，那样更快
                ans += max(a[k] - a[k - 1] - max(cnt0 * cnt0 - cnt1, 0), 0)
        return ans
```

```py [sol-Python3 更快写法]
class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        ans = tot1 = 0
        a = [-1]  # 哨兵
        for right, b in enumerate(s):
            if b == '0':
                a.append(right)
            else:
                ans += right - a[-1]  # 不含 0 子串的个数
                tot1 += 1
            for k in range(len(a) - 1, 0, -1):
                cnt0 = len(a) - k
                if cnt0 * cnt0 > tot1:
                    break
                cnt1 = right - a[k] + 1 - cnt0
                d = cnt0 * cnt0 - cnt1
                if d <= 0:
                    ans += a[k] - a[k - 1]
                elif (res := a[k] - a[k - 1] - d) > 0:
                    ans += res
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSubstrings(String S) {
        char[] s = S.toCharArray();
        int ans = 0;
        int tot1 = 0;
        int[] a = new int[s.length + 1];
        a[0] = -1; // 哨兵
        int m = 1;
        for (int right = 0; right < s.length; right++) {
            if (s[right] == '0') {
                a[m++] = right;
            } else {
                ans += right - a[m - 1];
                tot1++;
            }
            for (int k = m - 1; k > 0 && (m - k) * (m - k) <= tot1; k--) {
                int cnt0 = m - k;
                int cnt1 = right - a[k] + 1 - cnt0;
                ans += Math.max(a[k] - a[k - 1] - Math.max(cnt0 * cnt0 - cnt1, 0), 0);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSubstrings(string s) {
        int ans = 0, tot1 = 0;
        vector<int> a = {-1}; // 哨兵
        for (int right = 0; right < s.length(); right++) {
            char b = s[right];
            if (b == '0') {
                a.push_back(right);
            } else {
                ans += right - a.back();
                tot1++;
            }
            for (int k = a.size() - 1; k && (a.size() - k) * (a.size() - k) <= tot1; k--) {
                int cnt0 = a.size() - k;
                int cnt1 = right - a[k] + 1 - cnt0;
                ans += max(a[k] - a[k - 1] - max(cnt0 * cnt0 - cnt1, 0), 0);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubstrings(s string) (ans int) {
	tot1 := 0
	a := []int{-1} // 哨兵
	for right, b := range s {
		if b == '0' {
			a = append(a, right)
		} else {
			ans += right - a[len(a)-1]
			tot1++
		}
		for k := len(a) - 1; k > 0 && (len(a)-k)*(len(a)-k) <= tot1; k-- {
			cnt0 := len(a) - k
			cnt1 := right - a[k] + 1 - cnt0
			ans += max(a[k]-a[k-1]-max(cnt0*cnt0-cnt1, 0), 0)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{n})$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

注：使用队列，只保存 $\textit{right}$ 及其左侧的 $\mathcal{O}(\sqrt{n})$ 个 $0$ 的下标，可以把空间复杂度优化到 $\mathcal{O}(\sqrt{n})$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

## 方法一：二分答案

本题是标准的「**最小化最大值**」，用**二分答案**解决。如果你不知道什么是二分，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

> 为什么可以二分？因为子串长度越长，越能在 $\textit{numOps}$ 次操作内完成，有单调性。这意味着如果可以把每段子串的长度都变成至多为 $m$，那么也可以变成至多为 $m+1,m+2,\cdots$；如果不能把每段子串的长度都变成至多为 $m$，那么也不能变成至多为 $m-1,m-2,\cdots$。

二分长度的上界 $m$，问题变成：

- 判断能否在 $\textit{numOps}$ 次操作内，把每段连续相同子串的长度都变成 $\le m$ 的。

考虑其中一段连续相同子串要怎么改，最少要改多少次。

**核心想法**：好比切香肠，既然每一段至多为 $m$，那么就每隔 $m$ 个字符操作一次（切一刀）。

比如 $m=2$，举例说明：

- 子串长度为 $4$，例如 $\texttt{0000}\to \texttt{0010}$。改 $1$ 次。
- 子串长度为 $5$，例如 $\texttt{00000}\to \texttt{00100}$。改 $1$ 次。
- 子串长度为 $6$，例如 $\texttt{000000}\to \texttt{001010}$。改 $2$ 次。
- 子串长度为 $7$，例如 $\texttt{0000000}\to \texttt{0010010}$。改 $2$ 次。
- 子串长度为 $8$，例如 $\texttt{00000000}\to \texttt{00100100}$。改 $2$ 次。
- 子串长度为 $9$，例如 $\texttt{000000000}\to \texttt{001001010}$。改 $3$ 次。

一般地：

- 如果子串长度不是 $m+1$ 的倍数，那么可以修改子串中的下标为 $m,2m,3m,\cdots$ 的字符。
- 如果子串长度恰好是 $m+1$ 的倍数，那么最后一个字符不能改（否则会和下一段的字符一样），我们可以改子串的**倒数第二个**字符。

相当于每有 $m+1$ 个字符，就要操作一次。如果最后剩余字母不足 $m+1$，就不操作。所以修改次数为

$$
\left\lfloor\dfrac{子串长度}{m+1}\right\rfloor
$$

注意特判 $m=1$ 的情况，如果子串长度为 $2$，此时倒数第二个字符改完后，会把前一段的长度加一，可能导致不符合题目要求。又比如 $\texttt{0000}$，改成 $\texttt{0110}$ 不符合要求，改成 $\texttt{0101}$ 会让下一段的长度加一。所以不能按照上述公式计算，这种情况直接暴力判断改成 $\texttt{0101}\cdots$ 还是 $\texttt{1010}\cdots$，二者取最小值。

如果改成 $\texttt{0101}\cdots$，相当于把 $s[i]$ 改成 $i\bmod 2$，也就是比较 $s[i]$ 的 ASCII 值的奇偶性，是否等于 $i\bmod 2$ 的奇偶性。若不等，则计数器加一。也可以直接算出 $(s[i]\bmod 2)\oplus (i\bmod 2) = (s[i]\oplus i)\bmod 2$ 的值（其中 $\oplus$ 表示异或），即为计数器需要增加的量。

## 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。长度不能为 $0$，一定无法做到。
- 开区间右端点初始值：$n$。此时无需操作，一定可以做到。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

> 闭区间可以在 $[1,n-1]$ 中二分。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1wmkqYREnP/?t=18m08s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minLength(self, s: str, numOps: int) -> int:
        n = len(s)
        def check(m: int) -> bool:
            cnt = 0
            if m == 1:
                # 改成 0101...
                # 如果 s[i] 和 i 的奇偶性不同，cnt 加一
                cnt = sum((ord(b) ^ i) & 1 for i, b in enumerate(s))
                # n-cnt 表示改成 1010...
                cnt = min(cnt, n - cnt)
            else:
                k = 0
                for i, b in enumerate(s):
                    k += 1
                    # 到达连续相同子串的末尾
                    if i == n - 1 or b != s[i + 1]:
                        cnt += k // (m + 1)
                        k = 0
            return cnt <= numOps
        return bisect_left(range(n), True, 1, key=check)
```

```py [sol-Python3 groupby]
class Solution:
    def minLength(self, s: str, numOps: int) -> int:
        n = len(s)
        def check(m: int) -> bool:
            if m == 1:
                # 改成 0101...
                # 如果 s[i] 和 i 的奇偶性不同，cnt 加一
                cnt = sum((ord(b) ^ i) & 1 for i, b in enumerate(s))
                # n-cnt 表示改成 1010...
                cnt = min(cnt, n - cnt)
            else:
                cnt = sum(len(list(t)) // (m + 1) for _, t in groupby(s))
            return cnt <= numOps
        return bisect_left(range(n), True, 1, key=check)
```

```java [sol-Java]
class Solution {
    public int minLength(String S, int numOps) {
        char[] s = S.toCharArray();
        int left = 0;
        int right = s.length;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(mid, s, numOps)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int m, char[] s, int numOps) {
        int n = s.length;
        int cnt = 0;
        if (m == 1) {
            // 改成 0101...
            for (int i = 0; i < n; i++) {
                // 如果 s[i] 和 i 的奇偶性不同，cnt 加一
                cnt += (s[i] ^ i) & 1;
            }
            // n-cnt 表示改成 1010...
            cnt = Math.min(cnt, n - cnt);
        } else {
            int k = 0;
            for (int i = 0; i < n; i++) {
                k++;
                // 到达连续相同子串的末尾
                if (i == n - 1 || s[i] != s[i + 1]) {
                    cnt += k / (m + 1);
                    k = 0;
                }
            }
        }
        return cnt <= numOps;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(string s, int numOps) {
        int n = s.length();
        auto check = [&](int m) -> bool {
            int cnt = 0;
            if (m == 1) {
                // 改成 0101...
                for (int i = 0; i < n; i++) {
                    // 如果 s[i] 和 i 的奇偶性不同，cnt 加一
                    cnt += (s[i] ^ i) & 1;
                }
                // n-cnt 表示改成 1010...
                cnt = min(cnt, n - cnt);
            } else {
                int k = 0;
                for (int i = 0; i < n; i++) {
                    k++;
                    // 到达连续相同子串的末尾
                    if (i == n - 1 || s[i] != s[i + 1]) {
                        cnt += k / (m + 1);
                        k = 0;
                    }
                }
            }
            return cnt <= numOps;
        };
        int left = 0, right = n;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minLength(s string, numOps int) int {
	n := len(s)
	return 1 + sort.Search(n-1, func(m int) bool {
		m++
		cnt := 0
		if m == 1 {
			// 改成 0101...
			for i, b := range s {
				// 如果 s[i] 和 i 的奇偶性不同，cnt 加一
				cnt += (int(b) ^ i) & 1
			}
			// n-cnt 表示改成 1010...
			cnt = min(cnt, n-cnt)
		} else {
			k := 0
			for i := range n {
				k++
				// 到达连续相同子串的末尾
				if i == n-1 || s[i] != s[i+1] {
					cnt += k / (m + 1)
					k = 0
				}
			}
		}
		return cnt <= numOps
	})
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。二分 $\mathcal{O}(\log n)$ 次，每次 $\mathcal{O}(n)$ 遍历字符串 $s$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：最大堆

想一想，如果 $\textit{numOps}=1$，应该选哪个子串操作？

应该选最长的子串操作。如果选其他子串，那么最长子串长度不变，所以操作其他子串不会把答案变小。

如果 $\textit{numOps}=2$ 呢？操作完最长的，就操作第二长的。

更一般的情况呢？由于我们需要动态添加元素、查找最大元素和删除最大元素，所以用**最大堆**模拟。

堆中每个元素是一个三元组，包含：

1. **子串**操作后的最长**子段**长度。比如 $10$ 操作一次后变成 $5$ 和 $4$，其中最大值为 $5$。注意这里的「子段」是指子串分割后的「子子串」。
2. 原始子串长度。
3. 段数，初始为 $1$。

循环至多 $\textit{numOps}$ 次，每次循环：

1. 如果堆顶的最长子段长度等于 $2$，那么答案就是 $2$。注意需要提前特判答案为 $1$ 的情况，算法见方法一。
2. 否则，把原始子串**重新分割**。设原始子串长度为 $k$，当前段数为 $\textit{seg}$，那么操作后的最长子段长度为 $\left\lfloor\dfrac{k}{\textit{seg}+1}\right\rfloor$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1wmkqYREnP/?t=28m01s)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
class Solution:
    def minLength(self, s: str, numOps: int) -> int:
        cnt = sum((ord(b) ^ i) & 1 for i, b in enumerate(s))
        if min(cnt, len(s) - cnt) <= numOps:
            return 1

        g = (list(t) for _, t in groupby(s))
        # 子串操作后的最长子段长度，原始子串长度，段数
        h = [(-k, k, 1) for k in map(len, g)]
        heapify(h)
        for _ in range(numOps):
            max_seg, k, seg = h[0]
            if max_seg == -2:
                return 2
            heapreplace(h, (-(k // (seg + 1)), k, seg + 1))  # 重新分割
        return -h[0][0]
```

```java [sol-Java]
class Solution {
    public int minLength(String S, int numOps) {
        char[] s = S.toCharArray();
        int n = s.length;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt += (s[i] ^ i) & 1;
        }
        if (Math.min(cnt, n - cnt) <= numOps) {
            return 1;
        }

        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> (b[0] - a[0]));
        int k = 0;
        for (int i = 0; i < n; i++) {
            k++;
            // 到达连续相同子串的末尾
            if (i == n - 1 || s[i] != s[i + 1]) {
                pq.offer(new int[]{k, k, 1});
                k = 0;
            }
        }

        while (numOps-- > 0 && pq.peek()[0] > 2) {
            int[] t = pq.poll();
            t[0] = t[1] / ++t[2]; // 重新分割
            pq.offer(t);
        }
        return pq.peek()[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(string s, int numOps) {
        int n = s.length();
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt += (s[i] ^ i) & 1;
        }
        if (min(cnt, n - cnt) <= numOps) {
            return 1;
        }

        priority_queue<tuple<int, int, int>> pq;
        int k = 0;
        for (int i = 0; i < n; i++) {
            k++;
            // 到达连续相同子串的末尾
            if (i == n - 1 || s[i] != s[i + 1]) {
                pq.emplace(k, k, 1);
                k = 0;
            }
        }

        while (numOps-- && get<0>(pq.top()) > 2) {
            auto [_, k, seg] = pq.top();
            pq.pop();
            pq.emplace(k / (seg + 1), k, seg + 1); // 重新分割
        }
        return get<0>(pq.top());
    }
};
```

```go [sol-Go]
func minLength(s string, numOps int) int {
	n := len(s)
	cnt := 0
	for i, b := range s {
		cnt += (int(b) ^ i) & 1
	}
	if min(cnt, n-cnt) <= numOps {
		return 1
	}

	h := hp{}
	k := 0
	for i := 0; i < n; i++ {
		k++
		// 到达连续相同子串的末尾
		if i == n-1 || s[i] != s[i+1] {
			h = append(h, tuple{k, k, 1})
			k = 0
		}
	}
	heap.Init(&h)

	for ; numOps > 0 && h[0].maxSeg > 2; numOps-- {
		h[0].seg++
		h[0].maxSeg = h[0].k / h[0].seg // 重新分割
		heap.Fix(&h, 0)
	}
	return h[0].maxSeg
}

type tuple struct{ maxSeg, k, seg int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].maxSeg > h[j].maxSeg } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + \textit{numOps}\cdot \log n)$，其中 $n$ 是 $s$ 的长度。这里时间复杂度以 Python 和 Go 为主，初始化堆的时间是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

### 分桶优化

把堆中元素按照最长子段长度分组。每组是个 pair 列表，每个 pair 包含原始子串长度和段数。

这样就不需要堆了，用一个指针 $i$ 表示目前最长子段长度。

每次从第 $i$ 个组中弹出 pair $(k, \textit{seg})$，把 pair $(k, \textit{seg}+1)$ 加到第 $\left\lfloor\dfrac{k}{\textit{seg}+1}\right\rfloor$ 组中。

```py [sol-Python3]
class Solution:
    def minLength(self, s: str, numOps: int) -> int:
        n = len(s)
        cnt = sum((ord(b) ^ i) & 1 for i, b in enumerate(s))
        if min(cnt, n - cnt) <= numOps:
            return 1

        buckets = [[] for _ in range(n + 1)]
        for _, t in groupby(s):
            k = len(list(t))
            buckets[k].append((k, 1))  # 原始子串长度，段数

        i = n
        for _ in range(numOps):
            while not buckets[i]:
                i -= 1
            if i == 2:
                return 2
            k, seg = buckets[i].pop()
            buckets[k // (seg + 1)].append((k, seg + 1))

        while not buckets[i]:
            i -= 1
        return i
```

```java [sol-Java]
class Solution {
    public int minLength(String S, int numOps) {
        char[] s = S.toCharArray();
        int n = s.length;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt += (s[i] ^ i) & 1;
        }
        if (Math.min(cnt, n - cnt) <= numOps) {
            return 1;
        }

        List<int[]>[] buckets = new ArrayList[n + 1];
        Arrays.setAll(buckets, i -> new ArrayList<>());
        int k = 0;
        for (int i = 0; i < n; i++) {
            k++;
            // 到达连续相同子串的末尾
            if (i == n - 1 || s[i] != s[i + 1]) {
                buckets[k].add(new int[]{k, 1}); // 原始子串长度，段数
                k = 0;
            }
        }

        int i = n;
        while (numOps-- > 0) {
            while (buckets[i].isEmpty()) {
                i--;
            }
            if (i == 2) {
                return 2;
            }
            int[] p = buckets[i].remove(buckets[i].size() - 1);
            buckets[p[0] / ++p[1]].add(p);
        }

        while (buckets[i].isEmpty()) {
            i--;
        }
        return i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(string s, int numOps) {
        int n = s.length();
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt += (s[i] ^ i) & 1;
        }
        if (min(cnt, n - cnt) <= numOps) {
            return 1;
        }

        vector<vector<pair<int, int>>> buckets(n + 1);
        int k = 0;
        for (int i = 0; i < n; i++) {
            k++;
            // 到达连续相同子串的末尾
            if (i == n - 1 || s[i] != s[i + 1]) {
                buckets[k].emplace_back(k, 1); // 原始子串长度，段数
                k = 0;
            }
        }

        int i = n;
        while (numOps--) {
            while (buckets[i].empty()) {
                i--;
            }
            if (i == 2) {
                return 2;
            }
            auto [k, seg] = buckets[i].back();
            buckets[i].pop_back();
            buckets[k / (seg + 1)].emplace_back(k, seg + 1);
        }

        while (buckets[i].empty()) {
            i--;
        }
        return i;
    }
};
```

```go [sol-Go]
func minLength(s string, numOps int) int {
	n := len(s)
	cnt := 0
	for i, b := range s {
		cnt += (int(b) ^ i) & 1
	}
	if min(cnt, n-cnt) <= numOps {
		return 1
	}

	type pair struct{ k, seg int }
	h := make([][]pair, n+1)
	k := 0
	for i := 0; i < n; i++ {
		k++
		// 到达连续相同子串的末尾
		if i == n-1 || s[i] != s[i+1] {
			h[k] = append(h[k], pair{k, 1}) // 原始子串长度，段数
			k = 0
		}
	}

	i := n
	for range numOps {
		for len(h[i]) == 0 {
			i--
		}
		if i == 2 {
			return 2
		}
		p := h[i][len(h[i])-1]
		h[i] = h[i][:len(h[i])-1]
		p.seg++
		maxSeg := p.k / p.seg
		h[maxSeg] = append(h[maxSeg], p)
	}

	for len(h[i]) == 0 {
		i--
	}
	return i
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面二分题单中的「**最小化最大值**」，以及数据结构题单中的「**五、堆（优先队列）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

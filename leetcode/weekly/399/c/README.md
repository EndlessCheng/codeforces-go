## 方法一：统计因子

提示：

- 如果 $x$ 能被 $d$ 整除，那么 $x$ 必然有一个等于 $d$ 的因子。

思路：

为方便描述，把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 记作 $a$ 和 $b$。

1. 遍历 $a$，统计所有元素的因子个数，记录到哈希表 $\textit{cnt}$ 中。
2. 遍历 $b$，那么有 $\textit{cnt}[b[i]\cdot k]$ 个数可以被 $b[i]\cdot k$ 整除，加入答案。

优化：

1. 如果 $a[i]$ 不是 $k$ 的倍数，无法被 $b[i]\cdot k$ 整除，直接跳过不统计。
2. 此外，可以改为统计 $\dfrac{a[i]}{k}$ 的因子，这样需要循环的次数会更少；相应地，遍历 $b$ 时只需要把 $\textit{cnt}[b[i]]$ 加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, nums1: List[int], nums2: List[int], k: int) -> int:
        cnt = defaultdict(int)
        for x in nums1:
            if x % k:
                continue
            x //= k
            for d in range(1, isqrt(x) + 1):
                if x % d:
                    continue
                cnt[d] += 1
                if d * d < x:
                    cnt[x // d] += 1
        return sum(cnt[x] for x in nums2)
```

```java [sol-Java]
class Solution {
    public long numberOfPairs(int[] nums1, int[] nums2, int k) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums1) {
            if (x % k != 0) {
                continue;
            }
            x /= k;
            for (int d = 1; d * d <= x; d++) {
                if (x % d > 0) {
                    continue;
                }
                cnt.merge(d, 1, Integer::sum);
                if (d * d < x) {
                    cnt.merge(x / d, 1, Integer::sum);
                }
            }
        }

        long ans = 0;
        for (int x : nums2) {
            ans += cnt.getOrDefault(x, 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        unordered_map<int, int> cnt;
        for (int x : nums1) {
            if (x % k) {
                continue;
            }
            x /= k;
            for (int d = 1; d * d <= x; d++) {
                if (x % d) {
                    continue;
                }
                cnt[d]++;
                if (d * d < x) {
                    cnt[x / d]++;
                }
            }
        }

        long long ans = 0;
        for (int x : nums2) {
            ans += cnt.contains(x) ? cnt[x] : 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt := map[int]int{}
	for _, x := range nums1 {
		if x%k > 0 {
			continue
		}
		x /= k
		for d := 1; d*d <= x; d++ {
			if x%d == 0 {
				cnt[d]++
				if d*d < x {
					cnt[x/d]++
				}
			}
		}
	}

	for _, x := range nums2 {
		ans += int64(cnt[x])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{U/k} + m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。
- 空间复杂度：$\mathcal{O}(U/k)$。不同因子个数不会超过 $U/k$。

## 方法二：枚举倍数

统计 $a[i]/k$ 和 $b[i]$ 的出现次数，分别保存到哈希表 $\textit{cnt}_1$ 和 $\textit{cnt}_2$ 中。设 $\textit{cnt}_1$ 中的最大 key 为 $u$。

枚举 $\textit{cnt}_2$ 中的元素 $i$，然后枚举 $i$ 的倍数 $i,2i,3i,\cdots$，一直到 $u$，累加这些数在 $\textit{cnt}_1$ 中的 value，乘上 $\textit{cnt}_2[i]$，加入答案。

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, nums1: List[int], nums2: List[int], k: int) -> int:
        cnt1 = Counter(x // k for x in nums1 if x % k == 0)
        if not cnt1:
            return 0
        ans = 0
        u = max(cnt1)
        for i, c in Counter(nums2).items():
            s = sum(cnt1[j] for j in range(i, u + 1, i))
            ans += s * c
        return ans
```

```java [sol-Java]
public class Solution {
    public long numberOfPairs(int[] nums1, int[] nums2, int k) {
        Map<Integer, Integer> cnt1 = new HashMap<>();
        for (int x : nums1) {
            if (x % k == 0) {
                cnt1.merge(x / k, 1, Integer::sum);
            }
        }
        if (cnt1.isEmpty()) {
            return 0;
        }

        Map<Integer, Integer> cnt2 = new HashMap<>();
        for (int x : nums2) {
            cnt2.merge(x, 1, Integer::sum);
        }

        long ans = 0;
        int u = Collections.max(cnt1.keySet());
        for (Map.Entry<Integer, Integer> e : cnt2.entrySet()) {
            int s = 0;
            int i = e.getKey();
            for (int j = i; j <= u; j += i) {
                if (cnt1.containsKey(j)) {
                    s += cnt1.get(j);
                }
            }
            ans += (long) s * e.getValue();
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        unordered_map<int, int> cnt1;
        for (int x : nums1) {
            if (x % k == 0) {
                cnt1[x / k]++;
            }
        }
        if (cnt1.empty()) {
            return 0;
        }
        unordered_map<int, int> cnt2;
        for (int x : nums2) {
            cnt2[x]++;
        }

        long long ans = 0;
        int u = ranges::max_element(cnt1)->first;
        for (auto& [i, c] : cnt2) {
            int s = 0;
            for (int j = i; j <= u; j += i) {
                s += cnt1.contains(j) ? cnt1[j] : 0;
            }
            ans += (long long) s * c;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt1 := map[int]int{}
	for _, x := range nums1 {
		if x%k == 0 {
			cnt1[x/k]++
		}
	}
	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	u := slices.Max(nums1) / k
	for i, c := range cnt2 {
		s := 0
		for j := i; j <= u; j += i {
			s += cnt1[j]
		}
		ans += int64(s * c)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m + (U/k)\log m)$，其中 $n$ 是 $\textit{nums}_1$ 的长度，$m$ 是 $\textit{nums}_2$ 的长度，$U=\max(\textit{nums}_1)$。复杂度根据调和级数可得。详细解释请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/)。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

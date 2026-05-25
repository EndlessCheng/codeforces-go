设 $m$ 是 $\textit{nums}_1$ 的长度，$n$ 是 $\textit{nums}_2$ 的长度。注意 $m\le 5$。

暴力的想法是，用一个哈希表维护 $\textit{nums}_2$ 每个元素的出现次数，这样可以 $\mathcal{O}(m)$ 计算询问 2。但处理询问 1 需要 $\mathcal{O}(n)$ 时间，太慢了。

把 $\textit{nums}_2$ 分成若干段，每段 $B$ 个元素（最后一段可能不足 $B$ 个），一共有 $\left\lceil\dfrac{n}{B}\right\rceil$ 段。

对于每一段，用一个哈希表维护每个元素的出现次数，用一个变量 $\textit{add}$ 记录这一段整体增加了 $\textit{add}$（借鉴 Lazy 线段树的懒标记思想）。

- 询问 1：对于每一段，如果该段完全在 $[x,y]$ 中，把该段的 $\textit{add}$ 增加 $\textit{val}$；如果该段只有一部分在 $[x,y]$ 中，那么暴力更新在 $[x,y]$ 中的这部分的 $\textit{nums}_2$ 以及元素出现次数。时间复杂度为 $\mathcal{O}(B + n/B)$。
- 询问 2：遍历每一段以及 $\textit{nums}_1$，我们需要计算有多少个 $\textit{nums}_2[k]$ 满足 $\textit{nums}_1[j] + (\textit{nums}_2[k] + \textit{add}) = \textit{tot}$，即哈希表中的 $\textit{tot} - \textit{nums}_1[j] - \textit{add}$ 的出现次数。时间复杂度为 $\mathcal{O}(mn/B)$。

把初始化也算进来，总的时间复杂度为 

$$
\mathcal{O}(n + q(B+mn/B))
$$

根据基本不等式（或者对勾函数的性质），当 $B=\sqrt{mn}$ 时，上式取到最小值

$$
\mathcal{O}(n + q\sqrt{mn})
$$

[本题视频讲解](https://www.bilibili.com/video/BV16FG76JEQo/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, nums1: list[int], nums2: list[int], queries: list[list[int]]) -> list[int]:
        m, n = len(nums1), len(nums2)
        B = isqrt(m * n)

        # blocks[i] = [l, r, cnt, add]
        # l, r: 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
        # cnt: 这一段每个元素的出现次数
        # add: 这一段整体要增加 add
        blocks = []
        for i in range(0, n, B):
            r = min(i + B, n)
            blocks.append([i, r, Counter(nums2[i: r]), 0])

        ans = []
        for q in queries:
            if q[0] == 1:
                l, r, val = q[1], q[2] + 1, q[3]
                for i, (bl, br, cnt, _) in enumerate(blocks):
                    if br <= l:
                        continue
                    if bl >= r:
                        break
                    # blocks[i] 在 [l, r) 中
                    if l <= bl and br <= r:
                        blocks[i][3] += val
                        continue
                    # blocks[i] 的一部分在 [l, r) 中
                    # 暴力更新 nums2 的子数组的元素值及其出现次数
                    for j in range(max(bl, l), min(br, r)):
                        cnt[nums2[j]] -= 1  # 撤销旧的
                        nums2[j] += val
                        cnt[nums2[j]] += 1  # 添加新的
            else:
                res = 0
                for _, _, cnt, add in blocks:
                    target = q[1] - add
                    for x in nums1:
                        res += cnt.get(target - x, 0)  # 避免把不在 cnt 中的 key 插入哈希表
                ans.append(res)
        return ans
```

```java [sol-Java]
class Solution {
    public static final int MX = 1_000_000_001;

    // l, r: 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
    // cnt: 这一段每个元素的出现次数
    // add: 这一段整体要增加 add
    private record Block(int l, int r, Map<Integer, Integer> cnt, int add) {
    }

    public int[] numberOfPairs(int[] nums1, int[] nums2, int[][] queries) {
        int m = nums1.length;
        int n = nums2.length;
        int B = (int) Math.sqrt(m * n);

        Block[] blocks = new Block[(n - 1) / B + 1];
        for (int i = 0; i < n; i += B) {
            int r = Math.min(i + B, n);
            Map<Integer, Integer> cnt = new HashMap<>();
            for (int j = i; j < r; j++) {
                cnt.merge(nums2[j], 1, Integer::sum); // cnt[nums2[j]]++
            }
            blocks[i / B] = new Block(i, r, cnt, 0);
        }

        int cntQ2 = 0;
        for (int[] q : queries) {
            cntQ2 += q[0] - 1;
        }

        int[] ans = new int[cntQ2];
        int idx = 0;
        for (int[] q : queries) {
            if (q[0] == 1) {
                int l = q[1], r = q[2] + 1, val = q[3];
                for (int i = 0; i < blocks.length; i++) {
                    Block b = blocks[i];
                    if (b.l >= r) {
                        break;
                    }
                    if (b.r <= l || b.add >= MX) {
                        continue;
                    }
                    // b 在 [l, r) 中
                    if (l <= b.l && b.r <= r) {
                        blocks[i] = new Block(b.l, b.r, b.cnt, b.add + val);
                        continue;
                    }
                    // b 的一部分在 [l, r) 中
                    int bl = Math.max(b.l, l);
                    int br = Math.min(b.r, r);
                    // 暴力更新 nums2 的子数组 [bl, br) 的元素值及其出现次数
                    for (int j = bl; j < br; j++) {
                        b.cnt.merge(nums2[j], -1, Integer::sum); // 撤销旧的
                        nums2[j] = Math.min(nums2[j] + val, MX); // 避免溢出
                        b.cnt.merge(nums2[j], 1, Integer::sum); // 添加新的
                    }
                }
            } else {
                int tot = q[1];
                int res = 0;
                for (Block b : blocks) {
                    int target = tot - b.add;
                    for (int x : nums1) {
                        res += b.cnt.getOrDefault(target - x, 0);
                    }
                }
                ans[idx++] = res;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int MX = 1'000'000'001;

public:
    vector<int> numberOfPairs(vector<int>& nums1, vector<int>& nums2, vector<vector<int>>& queries) {
        int m = nums1.size(), n = nums2.size();
        int B = sqrt(m * n);

        // tuple<l, r, cnt, add>
        // l, r: 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
        // cnt: 这一段每个元素的出现次数
        // add: 这一段整体要增加 add
        vector<tuple<int, int, unordered_map<int, int>, int>> blocks;
        for (int i = 0; i < n; i += B) {
            int r = min(i + B, n);
            unordered_map<int, int> cnt;
            for (int j = i; j < r; j++) {
                cnt[nums2[j]]++;
            }
            blocks.emplace_back(i, r, cnt, 0);
        }

        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                int l = q[1], r = q[2] + 1, val = q[3];
                for (auto& [bl, br, cnt, add] : blocks) {
                    if (bl >= r) {
                        break;
                    }
                    if (br <= l || add >= MX) {
                        continue;
                    }
                    // b 在 [l, r) 中
                    if (l <= bl && br <= r) {
                        add += val;
                        continue;
                    }
                    // b 的一部分在 [l, r) 中
                    int L = max(bl, l);
                    int R = min(br, r);
                    // 暴力更新 nums2 的子数组 [L, R) 的元素值及其出现次数
                    for (int j = L; j < R; j++) {
                        cnt[nums2[j]]--; // 撤销旧的
                        nums2[j] = min(nums2[j] + val, MX); // 避免溢出
                        cnt[nums2[j]]++; // 添加新的
                    }
                }
            } else {
                int res = 0;
                for (auto& [_, _, cnt, add] : blocks) {
                    int target = q[1] - add;
                    for (int x : nums1) {
                        // 避免把 target - x 插入哈希表
                        auto it = cnt.find(target - x);
                        if (it != cnt.end()) {
                            res += it->second;
                        }
                    }
                }
                ans.push_back(res);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(nums1, nums2 []int, queries [][]int) (ans []int) {
	m, n := len(nums1), len(nums2)
	B := int(math.Sqrt(float64(m * n)))

	type block struct {
		l, r int         // 这一段对应 nums2 的子数组 [l, r)，注意是左闭右开区间
		cnt  map[int]int // 这一段每个元素的出现次数
		add  int         // 这一段整体要增加 add
	}
	blocks := make([]block, (n-1)/B+1)
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		cnt := map[int]int{}
		for _, x := range nums2[i:r] {
			cnt[x]++
		}
		blocks[i/B] = block{i, r, cnt, 0}
	}

	for _, q := range queries {
		if q[0] == 1 {
			l, r, val := q[1], q[2]+1, q[3]
			for i := range blocks {
				b := &blocks[i]
				if b.r <= l {
					continue
				}
				if b.l >= r {
					break
				}
				// b 在 [l, r) 中
				if l <= b.l && b.r <= r {
					b.add += val
					continue
				}
				// b 的一部分在 [l, r) 中
				bl := max(b.l, l)
				br := min(b.r, r)
				// 暴力更新 nums2 的子数组 [bl, br) 的元素值及其出现次数
				for j := bl; j < br; j++ {
					b.cnt[nums2[j]]-- // 撤销旧的
					nums2[j] += val
					b.cnt[nums2[j]]++ // 添加新的
				}
			}
		} else {
			res := 0
			for _, b := range blocks {
				target := q[1] - b.add
				for _, x := range nums1 {
					res += b.cnt[target-x]
				}
			}
			ans = append(ans, res)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\sqrt{mn})$，其中 $m$ 是 $\textit{nums}_1$ 的长度，$n$ 是 $\textit{nums}_2$ 的长度。注意 $m\le 5$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。所有哈希表的大小之和为 $\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**十、根号算法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

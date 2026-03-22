**前置知识**：[LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。

本题需要判断子数组是否包含 $\textit{or}$。我们可以记录 $\textit{nums}$ 每个元素的最近一次出现的位置 $\textit{last}[x]$，只要 $\textit{last}[\textit{or}]$ 大于等于子数组左端点，那么子数组就包含 $\textit{or}$。

设子数组右端点在 $i$，左端点在 $[\ell,r]$ 的子数组的 OR 都是 $\textit{or}$。设 $j = \textit{last}[\textit{or}]$。那么当 $j\ge \ell$ 时，左端点为 $\ell,\ell+1,\ldots,\min(r, j)$，右端点为 $i$ 的子数组都包含 $\textit{or}$，这一共有

$$
\min(r,j)-\ell+1
$$

个合法子数组，加入答案。

代码用到了原地去重算法，可以看 [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)，[我的题解](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solutions/2807162/gen-zhao-wo-guo-yi-bian-shi-li-2ni-jiu-m-rvyk/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countGoodSubarrays(self, nums: List[int]) -> int:
        or_left = []  # (子数组或值，最小左端点)
        last = {}
        ans = 0

        for i, x in enumerate(nums):
            last[x] = i

            # 计算以 i 为右端点的子数组或值
            for p in or_left:
                p[0] |= x
            # x 单独一个数作为子数组
            or_left.append([x, i])

            # 原地去重（相同或值只保留最左边的）
            # 原理见力扣 26. 删除有序数组中的重复项
            idx = 1
            for j in range(1, len(or_left)):
                if or_left[j][0] != or_left[j - 1][0]:
                    or_left[idx] = or_left[j]
                    idx += 1
            del or_left[idx:]

            for k, (or_val, left) in enumerate(or_left):
                right = or_left[k + 1][1] - 1 if k < len(or_left) - 1 else i
                # 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 or_val
                j = last.get(or_val, -1)
                if j >= left:
                    ans += min(right, j) - left + 1

        return ans
```

```java [sol-Java]
class Solution {
    public long countGoodSubarrays(int[] nums) {
        List<int[]> orLeft = new ArrayList<>(); // (子数组或值，最小左端点)
        Map<Integer, Integer> last = new HashMap<>();
        long ans = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            last.put(x, i);

            // 计算以 i 为右端点的子数组或值
            for (int[] p : orLeft) {
                p[0] |= x; // **根据题目修改**
            }
            // x 单独一个数作为子数组
            orLeft.add(new int[]{x, i});

            // 原地去重（相同或值只保留最左边的）
            // 原理见力扣 26. 删除有序数组中的重复项
            int m = 1;
            for (int j = 1; j < orLeft.size(); j++) {
                if (orLeft.get(j)[0] != orLeft.get(j - 1)[0]) {
                    orLeft.set(m++, orLeft.get(j));
                }
            }
            orLeft.subList(m, orLeft.size()).clear();

            for (int k = 0; k < m; k++) {
                int orVal = orLeft.get(k)[0];
                int left = orLeft.get(k)[1];
                int right = k < m - 1 ? orLeft.get(k + 1)[1] - 1 : i;
                // 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 orVal
                int j = last.getOrDefault(orVal, -1);
                if (j >= left) {
                    ans += Math.min(right, j) - left + 1;
                }
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countGoodSubarrays(vector<int>& nums) {
        vector<pair<int, int>> or_left; // (子数组或值，最小左端点)
        unordered_map<int, int> last;
        long long ans = 0;

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            last[x] = i;

            // 计算以 i 为右端点的子数组或值
            for (auto& [or_val, _] : or_left) {
                or_val |= x; // **根据题目修改**
            }
            // x 单独一个数作为子数组
            or_left.emplace_back(x, i);

            // 原地去重（相同或值只保留最左边的）
            // 原理见力扣 26. 删除有序数组中的重复项
            int m = 1;
            for (int j = 1; j < or_left.size(); j++) {
                if (or_left[j].first != or_left[j - 1].first) {
                    or_left[m++] = or_left[j];
                }
            }
            or_left.resize(m);

            for (int k = 0; k < m; k++) {
                auto [or_val, left] = or_left[k];
                int right = k + 1 < m ? or_left[k + 1].second - 1 : i;
                // 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 or_val
                auto it = last.find(or_val);
                if (it != last.end() && it->second >= left) {
                    ans += min(right, it->second) - left + 1;
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func countGoodSubarrays(nums []int) (ans int64) {
	type pair struct{ or, left int } // 子数组或值，最小左端点
	orLeft := []pair{}
	last := map[int]int{}

	for i, x := range nums {
		last[x] = i

		// 计算以 i 为右端点的子数组或值
		for j := range orLeft {
			orLeft[j].or |= x
		}
		// x 单独一个数作为子数组
		orLeft = append(orLeft, pair{x, i})

		// 原地去重（相同或值只保留最左边的）
		// 原理见力扣 26. 删除有序数组中的重复项
		idx := 1
		for j := 1; j < len(orLeft); j++ {
			if orLeft[j].or != orLeft[j-1].or {
				orLeft[idx] = orLeft[j]
				idx++
			}
		}
		orLeft = orLeft[:idx]

		for k, p := range orLeft {
			orVal := p.or
			left := p.left
			right := i
			if k < len(orLeft)-1 {
				right = orLeft[k+1].left - 1
			}
			// 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 orVal
			j, ok := last[orVal]
			if ok && j >= left {
				ans += int64(min(right, j) - left + 1)
			}
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log U)$。

## 专题训练

见下面位运算题单的「**AND/OR LogTrick**」。

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

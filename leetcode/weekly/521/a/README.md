## 方法一：暴力模拟

统计每个元素的出现次数 $\textit{cnt}$。

每轮循环，把 $\textit{cnt}$ 中出现次数大于 $0$ 的数，从小到大加入答案，然后把这些数的出现次数都减少一。

重复上述过程，直到答案的长度等于 $n$（$\textit{nums}$ 的长度）。

[本题视频讲解](https://www.bilibili.com/video/BV12gah6UE9b/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def rearrangeArray(self, nums: list[int]) -> list[int]:
        mx = max(nums)
        cnt = [0] * (mx + 1)
        for x in nums:
            cnt[x] += 1

        n = len(nums)
        ans = []
        while len(ans) < n:
            for x, c in enumerate(cnt):
                if c > 0:
                    ans.append(x)
                    cnt[x] -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] rearrangeArray(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        for (int x : nums) {
            cnt[x]++;
        }

        int n = nums.length;
        int[] ans = new int[n];
        int k = 0;
        while (k < n) {
            for (int x = 1; x <= mx; x++) {
                if (cnt[x] > 0) {
                    ans[k++] = x;
                    cnt[x]--;
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
    vector<int> rearrangeArray(vector<int>& nums) {
        int mx = ranges::max(nums);
        vector<int> cnt(mx + 1);
        for (int x : nums) {
            cnt[x]++;
        }

        int n = nums.size();
        vector<int> ans;
        while (ans.size() < n) {
            for (int x = 1; x <= mx; x++) {
                if (cnt[x] > 0) {
                    ans.push_back(x);
                    cnt[x]--;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func rearrangeArray(nums []int) []int {
	mx := slices.Max(nums)
	cnt := make([]int, mx+1)
	for _, x := range nums {
		cnt[x]++
	}

	n := len(nums)
	ans := make([]int, 0, n)
	for len(ans) < n {
		for x, c := range cnt {
			if c > 0 {
				ans = append(ans, x)
				cnt[x]--
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。返回值不计入。

## 方法二：按出现次数分组

示例 1 的 $\textit{nums}=[3,1,3,2,1,3]$，把 $1,2,3$ 分别视作三种正方体积木，同一种积木竖着叠起来，得到高度分别为 $2,1,3$ 的积木塔。

从下往上看：

- 第一层有积木 $1,2,3$。
- 第二层有积木 $1,3$。
- 第三层有积木 $3$。

把这些数串起来，就是答案 $[1,2,3] + [1,3] + [3] = [1,2,3,1,3,3]$。

```py [sol-Python3]
class Solution:
    def rearrangeArray(self, nums: list[int]) -> list[int]:
        mx = max(nums)
        cnt = [0] * (mx + 1)
        levels = [[] for _ in nums]
        for x in nums:
            c = cnt[x]
            levels[c].append(x)
            cnt[x] += 1

        ans = []
        # 从下到上遍历每一层的积木
        for level in levels:
            level.sort()
            ans += level
        return ans
```

```java [sol-Java]
class Solution {
    public int[] rearrangeArray(int[] nums) {
        int n = nums.length;
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        List<Integer>[] levels = new ArrayList[n];
        Arrays.setAll(levels, _ -> new ArrayList<>());

        for (int x : nums) {
            int c = cnt[x];
            levels[c].add(x);
            cnt[x]++;
        }

        int[] ans = new int[n];
        int k = 0;
        // 从下到上遍历每一层的积木
        for (List<Integer> level : levels) {
            level.sort(null);
            for (int x : level) {
                ans[k++] = x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> rearrangeArray(vector<int>& nums) {
        int mx = ranges::max(nums);
        vector<int> cnt(mx + 1);
        vector<vector<int>> levels(nums.size());
        for (int x : nums) {
            int c = cnt[x];
            levels[c].push_back(x);
            cnt[x]++;
        }

        vector<int> ans;
        // 从下到上遍历每一层的积木
        for (auto& level : levels) {
            ranges::sort(level);
            ans.insert(ans.end(), level.begin(), level.end());
        }
        return ans;
    }
};
```

```go [sol-Go]
func rearrangeArray(nums []int) []int {
	n := len(nums)
	mx := slices.Max(nums)
	cnt := make([]int, mx+1)
	levels := make([][]int, n)
	for _, x := range nums {
		c := cnt[x]
		levels[c] = append(levels[c], x)
		cnt[x]++
	}

	ans := make([]int, 0, n)
	// 从下到上遍历每一层的积木
	for _, level := range levels {
		slices.Sort(level)
		ans = append(ans, level...)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。瓶颈在排序上。创建大小为 $U$ 的数组需要 $\mathcal{O}(U)$ 的时间。如果改用哈希表，可以做到 $\mathcal{O}(n\log n)$ 时间。
- 空间复杂度：$\mathcal{O}(n + U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

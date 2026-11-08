## 问题分解

如果把 $\textit{nums}$ 的后缀 $[i,n-1]$ 递减排序后，$\textit{nums}$ 等于 $\textit{target}$ 的子数组 $[k,k+n-1]$（记作 $a$），那么必须满足：

1. **大前提**：$a$ 的每个数的出现次数恰好等于 $\textit{nums}$ 的每个数的出现次数。这是 [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)，可以用**滑动窗口**解决。[我的题解](https://leetcode.cn/problems/find-all-anagrams-in-a-string/solutions/2969498/liang-chong-fang-fa-ding-chang-hua-chuan-14pd/)。
2. $\textit{nums}$ 的前 $i$ 个数 $[0,i-1]$ 必须等于 $a$ 的前 $i$ 个数，我们需要知道 $i$ **最大是多少**。这可以用 **Z 函数**解决，构造数组 $A = \textit{nums} + \textit{target}$，然后计算 $A$ 的 Z 数组 $z$。
3. $a$ 的后 $n-i$ 个数必须是递减的（允许相邻相等），我们需要知道 $n-i$ 最大是多少，即 $i$ **最小是多少**。通过记录上一次相邻递增的位置，可以快速求出 $i$ 的最小值。

如果 1 和 2 都成立，那么 $a$ 的后 $n-i$ 个数的出现次数等于 $\textit{nums}$ 的后 $n-i$ 个数的出现次数，只要保证这两部分都是递减的，它们就一定相等。

这样就把原问题分成了三个问题，分别解决。

## 具体思路

先用 438 题找到 $\textit{target}$ 的所有合法子数组 $a$。不合法的 $\textit{ans}[k] = 0$。

根据 Z 数组的定义，数组 $a$ 的前 $\min(z[n+k],n)$ 个数与 $\textit{nums}$ 的相同，所以有

$$
i\le \min(z[n+k],n)
$$

设 $\textit{lastInc}[i]$ 是最大的满足 $j\le i$ 且 $\textit{target}[j-1] < \textit{target}[j]$ 下标 $j$。如果不存在则 $\textit{lastInc}[i] = 0$。

从子数组 $a$ 的右端点 $k+n-1$ 往左看，$k+i$ 必须 $\ge \textit{lastInc}[k+n-1]$，才能保证 $a$ 的后 $n-i$ 个数不存在递增的情况（即递减的），所以有

$$
i \ge \textit{lastInc}[k+n-1] - k
$$

此外，$i$ 必须满足 $0\le i\le n-1$。

综上所述，$i$ 必须满足

$$
\max(\textit{lastInc}[k+n-1] - k, 0) \le i \le \min(z[n+k],n-1)
$$

所以

$$
ans[k] = \max(\min(z[n+k],n-1) - \max(\textit{lastInc}[k+n-1] - k, 0) + 1, 0)
$$

注意最后要和 $0$ 取最大值，以保证结果非负。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    # 438. 找到字符串中所有字母异位词
    def findAnagrams(self, s: list[int], p: list[int]) -> list[bool]:
        diff = Counter(p)  # 统计 p 的每种整数的出现次数
        valid_l = [False] * (len(s) - len(p) + 1)
        left = 0
        for right, x in enumerate(s):
            diff[x] -= 1  # x 进入窗口
            while diff[x] < 0:  # x 太多了
                diff[s[left]] += 1  # 左端点离开窗口
                left += 1
            # 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
            if right - left + 1 == len(p):
                valid_l[left] = True
        return valid_l

    # 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
    def calcZ(self, s: list[int]) -> list[int]:
        n = len(s)
        z = [0] * n
        box_l = box_r = 0
        for i in range(1, n):
            if i <= box_r:
                z[i] = min(z[i - box_l], box_r - i + 1)
            while i + z[i] < n and s[z[i]] == s[i + z[i]]:
                box_l, box_r = i, i + z[i]
                z[i] += 1
        return z

    def countSuffix(self, nums: list[int], target: list[int]) -> list[int]:
        n, m = len(nums), len(target)
        valid_l = self.findAnagrams(target, nums)
        z = self.calcZ(nums + target)

        last_inc = [0] * m
        for i in range(1, m):
            if target[i - 1] < target[i]:
                last_inc[i] = i
            else:
                last_inc[i] = last_inc[i - 1]

        ans = [0] * len(valid_l)
        for k, valid in enumerate(valid_l):
            if valid:
                ans[k] = max(min(z[n + k], n - 1) - max(last_inc[k + n - 1] - k, 0) + 1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] countSuffix(int[] nums, int[] target) {
        boolean[] validL = findAnagrams(target, nums);

        int n = nums.length;
        int m = target.length;
        int[] s = Arrays.copyOf(nums, n + m);
        System.arraycopy(target, 0, s, n, m);
        int[] z = calcZ(s);

        int[] lastInc = new int[m];
        for (int i = 1; i < m; i++) {
            if (target[i - 1] < target[i]) {
                lastInc[i] = i;
            } else {
                lastInc[i] = lastInc[i - 1];
            }
        }

        int[] ans = new int[validL.length];
        for (int k = 0; k < validL.length; k++) {
            if (validL[k]) {
                ans[k] = Math.max(Math.min(z[n + k], n - 1) - Math.max(lastInc[k + n - 1] - k, 0) + 1, 0);
            }
        }
        return ans;
    }

    private boolean[] findAnagrams(int[] s, int[] p) {
        int mx = 0;
        for (int x : s) mx = Math.max(mx, x);
        for (int x : p) mx = Math.max(mx, x);

        int[] diff = new int[mx + 1]; // int[] 比 HashMap 快
        // 统计 p 的每种整数的出现次数
        for (int x : p) {
            diff[x]++;
        }

        boolean[] validL = new boolean[s.length - p.length + 1];
        int left = 0;
        for (int right = 0; right < s.length; right++) {
            int x = s[right];
            diff[x]--; // x 进入窗口
            while (diff[x] < 0) { // x 太多了
                diff[s[left]]++; // 左端点离开窗口
                left++;
            }
            if (right - left + 1 == p.length) { // 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
                validL[left] = true;
            }
        }
        return validL;
    }

    // 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
    private int[] calcZ(int[] s) {
        int n = s.length;
        int[] z = new int[n];
        int boxL = 0;
        int boxR = 0;
        for (int i = 1; i < n; i++) {
            if (i <= boxR) {
                z[i] = Math.min(z[i - boxL], boxR - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                boxL = i;
                boxR = i + z[i];
                z[i]++;
            }
        }
        return z;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 438. 找到字符串中所有字母异位词
    vector<int8_t> findAnagrams(vector<int>& s, vector<int>& p) {
        int mx = max(ranges::max(s), ranges::max(p));
        vector<int> diff(mx + 1); // vector 比 unordered_map 快
        // 统计 p 的每种整数的出现次数
        for (int x : p) {
            diff[x]++;
        }

        vector<int8_t> valid_l(s.size() - p.size() + 1);
        int left = 0;
        for (int right = 0; right < s.size(); right++) {
            int x = s[right];
            diff[x]--; // x 进入窗口
            while (diff[x] < 0) { // x 太多了
                diff[s[left]]++; // 左端点离开窗口
                left++;
            }
            if (right - left + 1 == p.size()) { // 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
                valid_l[left] = true;
            }
        }
        return valid_l;
    }

    // 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
    vector<int> calcZ(vector<int>& s) {
        int n = s.size();
        vector<int> z(n);
        int box_l = 0, box_r = 0;
        for (int i = 1; i < n; i++) {
            if (i <= box_r) {
                z[i] = min(z[i - box_l], box_r - i + 1);
            }
            while (i + z[i] < n && s[z[i]] == s[i + z[i]]) {
                box_l = i;
                box_r = i + z[i];
                z[i]++;
            }
        }
        return z;
    }

public:
    vector<int> countSuffix(vector<int>& nums, vector<int>& target) {
        int n = nums.size(), m = target.size();
        vector<int8_t> valid_l = findAnagrams(target, nums);
        nums.insert(nums.end(), target.begin(), target.end());
        vector<int> z = calcZ(nums);

        vector<int> last_inc(m);
        for (int i = 1; i < m; i++) {
            if (target[i - 1] < target[i]) {
                last_inc[i] = i;
            } else {
                last_inc[i] = last_inc[i - 1];
            }
        }

        vector<int> ans(valid_l.size());
        for (int k = 0; k < valid_l.size(); k++) {
            if (valid_l[k]) {
                ans[k] = max(min(z[n + k], n - 1) - max(last_inc[k + n - 1] - k, 0) + 1, 0);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 438. 找到字符串中所有字母异位词
func findAnagrams(s, p []int) []bool {
	mx := max(slices.Max(s), slices.Max(p))
	diff := make([]int, mx+1) // slice 比 map 快
	// 统计 p 的每种整数的出现次数
	for _, x := range p {
		diff[x]++
	}

	validL := make([]bool, len(s)-len(p)+1)
	left := 0
	for right, x := range s {
		diff[x]-- // x 进入窗口
		for diff[x] < 0 { // x 太多了
			diff[s[left]]++ // 左端点离开窗口
			left++
		}
		if right-left+1 == len(p) { // 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
			validL[left] = true
		}
	}
	return validL
}

// 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
func calcZ(s []int) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func countSuffix(nums []int, target []int) []int {
	n, m := len(nums), len(target)
	validL := findAnagrams(target, nums)
	z := calcZ(append(nums, target...))

	lastInc := make([]int, m)
	for i := 1; i < m; i++ {
		if target[i-1] < target[i] {
			lastInc[i] = i
		} else {
			lastInc[i] = lastInc[i-1]
		}
	}

	ans := make([]int, len(validL))
	for k, valid := range validL {
		if valid {
			ans[k] = max(min(z[n+k], n-1)-max(lastInc[k+n-1]-k, 0)+1, 0)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$ 或 $\mathcal{O}(n+m+U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是 $\textit{target}$ 的长度，$U=\max(\max(\textit{nums}),\max(\textit{target}))$。
- 空间复杂度：$\mathcal{O}(n+m)$ 或 $\mathcal{O}(n+m+U)$。

## 专题训练

1. 滑动窗口题单。
2. 字符串题单的「**二、Z 函数**」。

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

如果所有 $\textit{nums}[i]$ 都是 $1$，那么子数组和等于子数组长度，问题变成 [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)。这可以 **Manacher 算法** 解决，原理见 [视频讲解](https://www.bilibili.com/video/BV1UcyYY4EnQ/)，欢迎点赞关注~

本题 $\textit{nums}$ 中的数都是非负数，所以子数组越长越好。

用 Manacher 算法，可以求出以 $\textit{nums}[i]$ 为回文中心（或者以 $\textit{nums}[i]$ 和 $\textit{nums}[i+1]$ 为回文中心）的最长回文子数组的范围。

枚举回文中心，用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 计算子数组和。

```py [sol-Python3]
class Solution:
    def getSum(self, s: List[int]) -> int:
        # Manacher 模板
        # 将 s 改造为 t，这样就不需要讨论 len(s) 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        # s 和 t 的下标转换关系：
        # (si+1)*2 = ti
        # ti/2-1 = si
        # ti 为偶数，对应奇回文串（从 2 开始）
        # ti 为奇数，对应偶回文串（从 3 开始）
        t = [-2, -1]
        for x in s:
            t.append(-1)
            t.append(x)
        t.append(-1)
        t.append(-3)

        # 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        # half_len[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        # 即 [i-half_len[i]+1,i+half_len[i]-1] 是 t 上的一个回文子串
        half_len = [0] * (len(t) - 2)
        half_len[1] = 1
        # box_r 表示当前右边界下标最大的回文子串的右边界下标+1
        # box_m 为该回文子串的中心位置，二者的关系为 r=mid+half_len[mid]
        box_m = box_r = max_i = 0
        for i in range(2, len(half_len)):
            hl = 1
            if i < box_r:
                # 记 i 关于 box_m 的对称位置 i'=box_m*2-i
                # 若以 i' 为中心的最长回文子串范围超出了以 box_m 为中心的回文串的范围（即 i+half_len[i'] >= box_r）
                # 则 half_len[i] 应先初始化为已知的回文半径 box_r-i，然后再继续暴力匹配
                # 否则 half_len[i] 与 half_len[i'] 相等
                hl = min(half_len[box_m * 2 - i], box_r - i)

            # 暴力扩展
            # 算法的复杂度取决于这部分执行的次数
            # 由于扩展之后 box_r 必然会更新（右移），且扩展的的次数就是 box_r 右移的次数
            # 因此算法的复杂度 = O(len(t)) = O(n)
            while t[i - hl] == t[i + hl]:
                hl += 1
                box_m, box_r = i, i + hl

            half_len[i] = hl
            if hl > half_len[max_i]:
                max_i = i

        pre = list(accumulate(s, initial=0))
        ans = 0

        for i in range(2, len(half_len)):
            hl = half_len[i]
            # 注意 t 上的最长回文子串的最左边和最右边都是 -1
            # 所以要对应到 s，最长回文子串的下标是从 i-hl+2 到 i+hl-2
            # 结合上文的下标转换关系，得到其在 s 上的下标范围是从 (i-hl)/2 到 (i+hl)/2-2
            ans = max(ans, pre[(i + hl) // 2 - 1] - pre[(i - hl) // 2])
        return ans
```

```java [sol-Java]
class Solution {
    public long getSum(int[] s) {
        // Manacher 模板
        // 将 s 改造为 t，这样就不需要讨论 s.length 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        // s 和 t 的下标转换关系：
        // (si+1)*2 = ti
        // ti/2-1 = si
        // ti 为偶数，对应奇回文串（从 2 开始）
        // ti 为奇数，对应偶回文串（从 3 开始）
        int n = s.length;
        int[] t = new int[n * 2 + 3];
        Arrays.fill(t, -1);
        t[0] = -2;
        for (int i = 0; i < n; i++) {
            t[i * 2 + 2] = s[i];
        }
        t[n * 2 + 2] = -3;

        // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        // halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        // 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
        int[] halfLen = new int[t.length - 2];
        halfLen[1] = 1;

        // maxI 记录最长回文子串在 halfLen 中的下标
        int maxI = 0;
        // boxR 表示当前右边界下标最大的回文子串的右边界下标+1
        // boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
        int boxM = 0;
        int boxR = 0;
        for (int i = 2; i < halfLen.length; i++) {
            int hl = 1;
            if (i < boxR) {
                // 记 i 关于 boxM 的对称位置 i'=boxM*2-i
                // 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
                // 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
                // 否则 halfLen[i] 与 halfLen[i'] 相等
                hl = Math.min(halfLen[boxM * 2 - i], boxR - i);
            }

            // 暴力扩展
            while (t[i - hl] == t[i + hl]) {
                hl++;
                boxM = i;
                boxR = i + hl;
            }

            halfLen[i] = hl;
            if (hl > halfLen[maxI]) {
                maxI = i;
            }
        }

        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + s[i];
        }

        long ans = 0;
        for (int i = 2; i < halfLen.length; i++) {
            int hl = halfLen[i];
            // 注意 t 上的最长回文子串的最左边和最右边都是 -1
            // 所以要对应到 s，最长回文子串的下标是从 i-hl+2 到 i+hl-2
            // 结合上文的下标转换关系，得到其在 s 上的下标范围是从 (i-hl)/2 到 (i+hl)/2-2
            ans = Math.max(ans, sum[(i + hl) / 2 - 1] - sum[(i - hl) / 2]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long getSum(vector<int>& s) {
        // Manacher 模板
        // 将 s 改造为 t，这样就不需要讨论 s.size() 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        // s 和 t 的下标转换关系：
        // (si+1)*2 = ti
        // ti/2-1 = si
        // ti 为偶数，对应奇回文串（从 2 开始）
        // ti 为奇数，对应偶回文串（从 3 开始）
        vector<int> t;
        t.reserve(s.size() * 2 + 3);
        t.push_back(-2);
        for (int x : s) {
            t.push_back(-1);
            t.push_back(x);
        }
        t.push_back(-1);
        t.push_back(-3);

        // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        // half_len[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        // 即 [i-half_len[i]+1,i+half_len[i]-1] 是 t 上的一个回文子串
        vector<int> half_len(t.size() - 2);
        half_len[1] = 1;
        // box_r 表示当前右边界下标最大的回文子串的右边界下标+1
        // box_m 为该回文子串的中心位置，二者的关系为 r=mid+half_len[mid]
        int box_m = 0, box_r = 0, max_i = 0;
        for (int i = 2; i < half_len.size(); i++) {
            int hl = 1;
            if (i < box_r) {
                // 记 i 关于 box_m 的对称位置 i'=box_m*2-i
                // 若以 i' 为中心的最长回文子串范围超出了以 box_m 为中心的回文串的范围（即 i+half_len[i'] >= box_r）
                // 则 half_len[i] 应先初始化为已知的回文半径 box_r-i，然后再继续暴力匹配
                // 否则 half_len[i] 与 half_len[i'] 相等
                hl = min(half_len[box_m * 2 - i], box_r - i);
            }

            // 暴力扩展
            // 算法的复杂度取决于这部分执行的次数
            // 由于扩展之后 box_r 必然会更新（右移），且扩展的的次数就是 box_r 右移的次数
            // 因此算法的复杂度 = O(t.size()) = O(n)
            while (t[i - hl] == t[i + hl]) {
                hl++;
                box_m = i;
                box_r = i + hl;
            }

            half_len[i] = hl;
            if (hl > half_len[max_i]) {
                max_i = i;
            }
        }

        vector<long long> sum(s.size() + 1);
        for (int i = 0; i < s.size(); i++) {
            sum[i + 1] = sum[i] + s[i];
        }

        long long ans = 0;
        for (int i = 2; i < half_len.size(); i++) {
            int hl = half_len[i];
            // 注意 t 上的最长回文子串的最左边和最右边都是 -1
            // 所以要对应到 s，最长回文子串的下标是从 i-hl+2 到 i+hl-2
            // 结合上文的下标转换关系，得到其在 s 上的下标范围是从 (i-hl)/2 到 (i+hl)/2-2
            ans = max(ans, sum[(i + hl) / 2 - 1] - sum[(i - hl) / 2]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func getSum(s []int) (ans int64) {
	// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新数组 t 的每个回文子数组都是奇回文子数组（都有回文中心）
	// s 和 t 的下标转换关系：
	// (si+1)*2 = ti
	// ti/2-1 = si
	// ti 为偶数（2,4,6,...）对应 s 中的奇回文子数组
	// ti 为奇数（3,5,7,...）对应 s 中的偶回文子数组
	t := append(make([]int, 0, len(s)*2+3), -2)
	for _, c := range s {
		t = append(t, -1, c)
	}
	t = append(t, -1, -3)

	// 定义一个奇回文子数组的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余子数组的长度
	// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子数组的回文半径
	// 具体地，闭区间 [i-halfLen[i]+1, i+halfLen[i]-1] 是 t 上的一个回文子数组
	// 由于 t 中回文子数组的首尾元素一定是 -1，根据下标转换关系，
	// 可以得到其在 s 中对应的回文子数组的区间为 [(i-halfLen[i])/2, (i+halfLen[i])/2-2]，用这个结论去计算子数组和
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	// boxR 表示当前右边界下标最大的回文子数组的右边界下标+1（初始化成任意 <= 0 的数都可以）
	// boxM 为该最大回文子数组的中心位置，二者的关系为 boxR = boxM + halfLen[boxM]
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原数组的首尾元素
		hl := 1
		if i < boxR {
			// 记 i 关于 boxM 的对称位置 i'=boxM*2-i
			// 若以 i' 为中心的最长回文子数组范围超出了以 boxM 为中心的回文子数组的范围（即 i+halfLen[i'] >= boxR）
			// 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
			// 否则 halfLen[i] 与 halfLen[i'] 相等
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		// 暴力扩展
		// 算法的复杂度取决于这部分执行的次数
		// 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
		// 因此算法的复杂度 = O(len(t)) = O(len(s))
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
	}

	sum := make([]int64, len(s)+1)
	for i, x := range s {
		sum[i+1] = sum[i] + int64(x)
	}

	for i := 2; i < len(halfLen); i++ {
		hl := halfLen[i]
		// 见上面注释
		ans = max(ans, sum[(i+hl)/2-1]-sum[(i-hl)/2])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面字符串题单的「**三、Manacher 算法**」。

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

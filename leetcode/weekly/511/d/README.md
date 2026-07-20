根据题意，两个字符串等价，意思是：

- 偶数下标的子序列循环同构。
- 奇数下标的子序列循环同构。

为了判断多个字符串是否等价，我们可以把每个字符串的偶数下标子序列 $a$ 变成其「最小表示」：在 $a$ 的所有循环同构字符串中，字典序最小的字符串。奇数下标子序列同理。然后把得到的字符串插入哈希集合中。最后，哈希集合的大小，即为最小组数。

如何计算字符串的最小表示？原理见视频讲解和代码注释。

[本题视频讲解](https://www.bilibili.com/video/BV1xpK663Eqh/)，欢迎点赞关注~

```py [sol-Python3]
# 返回 s 的字典序最小的循环同构串
# 时间复杂度 O(|s|)，证明见代码末尾的注释
def smallestRepresentation(s: str) -> str:
    n = len(s)
    s += s
    i = 0  # 始终指向当前最小子串的首字母下标
    j = 1  # 指向需要和 i 比较的子串的首字母下标
    while j < n:
        # 暴力比较：是 i 开头的字典序小，还是 j 开头的字典序小？
        k = 0
        while k < n and s[i + k] == s[j + k]:
            k += 1
        if k >= n:
            # s 是个周期字符串，周期为 j-i
            # j+d 开头的子串等于 i+d 开头的子串，而这些子串我们之前已经排除了，继续遍历不会找到更小的
            break

        if s[i + k] < s[j + k]:  # 注：如果求字典序最大，改成 >
            # 比如从 i 开始是 "aaab"，从 j 开始是 "aaac"
            # 从 i 开始比从 j 开始更小（排除 j）
            # 此外：
            # 从 i+1 开始比从 j+1 开始更小，所以从 j+1 开始不可能是答案，排除
            # 从 i+2 开始比从 j+2 开始更小，所以从 j+2 开始不可能是答案，排除
            # ……
            # 从 i+k 开始比从 j+k 开始更小，所以从 j+k 开始不可能是答案，排除
            # 所以下一个「可能是答案」的开始位置是 j+k+1
            j += k + 1
        else:
            # 从 j 开始比从 i 开始更小，更新 i=j（也意味着我们排除了 i）
            # 此外：
            # 从 j+1 开始比从 i+1 开始更小，所以从 i+1 开始不可能是答案，排除
            # 从 j+2 开始比从 i+2 开始更小，所以从 i+2 开始不可能是答案，排除
            # ……
            # 从 j+k 开始比从 i+k 开始更小，所以从 i+k 开始不可能是答案，排除
            # 所以把 j 跳到 i+k+1，不过这可能比 j+1 小，所以与 j+1 取 max
            # 综上所述，下一个「可能是答案」的开始位置是 max(j+1, i+k+1)
            i, j = j, max(j, i + k) + 1

        # 每次要么排除 k+1 个与 i 相关的位置（这样的位置至多 n 个），要么排除 k+1 个与 j 相关的位置（这样的位置至多 n 个）
        # 所以上面关于 k 的循环，∑k <= 2n，所以二重循环的总循环次数是 O(n) 的

    return s[i: i + n]


class Solution:
    def minimumGroups(self, words: list[str]) -> int:
        st = set()

        for word in words:
            min_s = [''] * len(word)

            # 偶数下标
            s = smallestRepresentation(word[::2])
            for j, ch in enumerate(s):
                min_s[j * 2] = ch

            # 奇数下标
            s = smallestRepresentation(word[1::2])
            for j, ch in enumerate(s):
                min_s[j * 2 + 1] = ch

            st.add(''.join(min_s))

        return len(st)
```

```java [sol-Java]
class Solution {
    public int minimumGroups(String[] words) {
        Set<String> st = new HashSet<>();

        for (String word : words) {
            char[] w = word.toCharArray();
            int m = w.length;
            char[] minS = new char[m];

            // 偶数下标
            char[] even = new char[(m + 1) / 2];
            for (int i = 0; i < even.length; i++) {
                even[i] = w[i * 2];
            }
            char[] s = smallestRepresentation(even);
            for (int j = 0; j < s.length; j++) {
                minS[j * 2] = s[j];
            }

            // 奇数下标
            char[] odd = new char[m / 2];
            for (int i = 0; i < odd.length; i++) {
                odd[i] = w[i * 2 + 1];
            }
            s = smallestRepresentation(odd);
            for (int j = 0; j < s.length; j++) {
                minS[j * 2 + 1] = s[j];
            }

            st.add(new String(minS));
        }

        return st.size();
    }

    // 返回 str 的字典序最小的循环同构串
    // 时间复杂度 O(|str|)，证明见代码末尾的注释
    private char[] smallestRepresentation(char[] str) {
        int n = str.length;
        char[] s = new char[n * 2]; // s = str + str
        System.arraycopy(str, 0, s, 0, n);
        System.arraycopy(str, 0, s, n, n);

        int i = 0; // 始终指向当前最小子串的首字母下标
        int j = 1; // 指向需要和 i 比较的子串的首字母下标
        while (j < n) {
            // 暴力比较：是 i 开头的字典序小，还是 j 开头的字典序小？
            int k = 0;
            while (k < n && s[i + k] == s[j + k]) {
                k++;
            }
            if (k >= n) {
                // s 是个周期字符串，周期为 j-i
                // j+d 开头的子串等于 i+d 开头的子串，而这些子串我们之前已经排除了，继续遍历不会找到更小的
                break;
            }

            if (s[i + k] < s[j + k]) { // 注：如果求字典序最大，改成 >
                // 比如从 i 开始是 "aaab"，从 j 开始是 "aaac"
                // 从 i 开始比从 j 开始更小（排除 j）
                // 此外：
                // 从 i+1 开始比从 j+1 开始更小，所以从 j+1 开始不可能是答案，排除
                // 从 i+2 开始比从 j+2 开始更小，所以从 j+2 开始不可能是答案，排除
                // ……
                // 从 i+k 开始比从 j+k 开始更小，所以从 j+k 开始不可能是答案，排除
                // 所以下一个「可能是答案」的开始位置是 j+k+1
                j += k + 1;
            } else {
                // 从 j 开始比从 i 开始更小，更新 i=j（也意味着我们排除了 i）
                // 此外：
                // 从 j+1 开始比从 i+1 开始更小，所以从 i+1 开始不可能是答案，排除
                // 从 j+2 开始比从 i+2 开始更小，所以从 i+2 开始不可能是答案，排除
                // ……
                // 从 j+k 开始比从 i+k 开始更小，所以从 i+k 开始不可能是答案，排除
                // 所以把 j 跳到 i+k+1，不过这可能比 j+1 小，所以与 j+1 取 max
                // 综上所述，下一个「可能是答案」的开始位置是 max(j+1, i+k+1)
                int tmp = j;
                j = Math.max(j, i + k) + 1;
                i = tmp;
            }

            // 每次要么排除 k+1 个与 i 相关的位置（这样的位置至多 n 个），要么排除 k+1 个与 j 相关的位置（这样的位置至多 n 个）
            // 所以上面关于 k 的循环，∑k <= 2n，所以二重循环的总循环次数是 O(n) 的
        }
        return Arrays.copyOfRange(s, i, i + n);
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回 s 的字典序最小的循环同构串
    // 时间复杂度 O(|s|)，证明见代码末尾的注释
    string smallestRepresentation(string& s) {
        int n = s.size();
        s += s;
        int i = 0; // 始终指向当前最小子串的首字母下标
        int j = 1; // 指向需要和 i 比较的子串的首字母下标
        while (j < n) {
            // 暴力比较：是 i 开头的字典序小，还是 j 开头的字典序小？
            int k = 0;
            while (k < n && s[i + k] == s[j + k]) {
                k++;
            }
            if (k >= n) {
                // s 是个周期字符串，周期为 j-i
                // j+d 开头的子串等于 i+d 开头的子串，而这些子串我们之前已经排除了，继续遍历不会找到更小的
                break;
            }

            if (s[i + k] < s[j + k]) { // 注：如果求字典序最大，改成 >
                // 比如从 i 开始是 "aaab"，从 j 开始是 "aaac"
                // 从 i 开始比从 j 开始更小（排除 j）
                // 此外：
                // 从 i+1 开始比从 j+1 开始更小，所以从 j+1 开始不可能是答案，排除
                // 从 i+2 开始比从 j+2 开始更小，所以从 j+2 开始不可能是答案，排除
                // ……
                // 从 i+k 开始比从 j+k 开始更小，所以从 j+k 开始不可能是答案，排除
                // 所以下一个「可能是答案」的开始位置是 j+k+1
                j += k + 1;
            } else {
                // 从 j 开始比从 i 开始更小，更新 i=j（也意味着我们排除了 i）
                // 此外：
                // 从 j+1 开始比从 i+1 开始更小，所以从 i+1 开始不可能是答案，排除
                // 从 j+2 开始比从 i+2 开始更小，所以从 i+2 开始不可能是答案，排除
                // ……
                // 从 j+k 开始比从 i+k 开始更小，所以从 i+k 开始不可能是答案，排除
                // 所以把 j 跳到 i+k+1，不过这可能比 j+1 小，所以与 j+1 取 max
                // 综上所述，下一个「可能是答案」的开始位置是 max(j+1, i+k+1)
                int tmp = j;
                j = max(j, i + k) + 1;
                i = tmp;
            }

            // 每次要么排除 k+1 个与 i 相关的位置（这样的位置至多 n 个），要么排除 k+1 个与 j 相关的位置（这样的位置至多 n 个）
            // 所以上面关于 k 的循环，∑k <= 2n，所以二重循环的总循环次数是 O(n) 的
        }
        return s.substr(i, n);
    }

public:
    int minimumGroups(vector<string>& words) {
        unordered_set<string> st;

        for (auto& word : words) {
            // 按照下标的奇偶性分组
            string groups[2]{};
            for (int i = 0; i < word.size(); i++) {
                groups[i % 2] += word[i];
            }

            // 分别计算偶数组和奇数组的最小表示
            string min_s(word.size(), 0);
            for (int k = 0; k < 2; k++) {
                auto s = smallestRepresentation(groups[k]);
                for (int j = 0; j < s.size(); j++) {
                    min_s[j * 2 + k] = s[j];
                }
            }

            st.insert(min_s);
        }

        return st.size();
    }
};
```

```go [sol-Go]
// 返回 s 的字典序最小的循环同构串
// 时间复杂度 O(|s|)，证明见代码末尾的注释
func smallestRepresentation(s []byte) []byte {
	n := len(s)
	s = append(s, s...)
	i := 0 // 始终指向当前最小子串的首字母下标
	j := 1 // 指向需要和 i 比较的子串的首字母下标
	for j < n {
		// 暴力比较：是 i 开头的字典序小，还是 j 开头的字典序小？
		k := 0
		for k < n && s[i+k] == s[j+k] {
			k++
		}
		if k >= n {
			// s 是个周期字符串，周期为 j-i
			// j+d 开头的子串等于 i+d 开头的子串，而这些子串我们之前已经排除了，继续遍历不会找到更小的
			break
		}

		if s[i+k] < s[j+k] { // 注：如果求字典序最大，改成 >
			// 比如从 i 开始是 "aaab"，从 j 开始是 "aaac"
			// 从 i 开始比从 j 开始更小（排除 j）
			// 此外：
			// 从 i+1 开始比从 j+1 开始更小，所以从 j+1 开始不可能是答案，排除
			// 从 i+2 开始比从 j+2 开始更小，所以从 j+2 开始不可能是答案，排除
			// ……
			// 从 i+k 开始比从 j+k 开始更小，所以从 j+k 开始不可能是答案，排除
			// 所以下一个「可能是答案」的开始位置是 j+k+1
			j += k + 1
		} else {
			// 从 j 开始比从 i 开始更小，更新 i=j（也意味着我们排除了 i）
			// 此外：
			// 从 j+1 开始比从 i+1 开始更小，所以从 i+1 开始不可能是答案，排除
			// 从 j+2 开始比从 i+2 开始更小，所以从 i+2 开始不可能是答案，排除
			// ……
			// 从 j+k 开始比从 i+k 开始更小，所以从 i+k 开始不可能是答案，排除
			// 所以把 j 跳到 i+k+1，不过这可能比 j+1 小，所以与 j+1 取 max
			// 综上所述，下一个「可能是答案」的开始位置是 max(j+1, i+k+1)
			i, j = j, max(j, i+k)+1
		}
		// 每次要么排除 k+1 个与 i 相关的位置（这样的位置至多 n 个），要么排除 k+1 个与 j 相关的位置（这样的位置至多 n 个）
		// 所以上面关于 k 的循环，∑k <= 2n，所以二重循环的总循环次数是 O(n) 的
	}
	return s[i : i+n]
}

func minimumGroups(words []string) (ans int) {
	set := map[string]struct{}{}

	for _, word := range words {
		// 按照下标的奇偶性分组
		groups := [2][]byte{}
		for i, ch := range word {
			groups[i%2] = append(groups[i%2], byte(ch))
		}

		// 分别计算偶数组和奇数组的最小表示
		minS := make([]byte, len(word))
		for k, s := range groups {
			s = smallestRepresentation(s)
			for j, ch := range s {
				minS[j*2+k] = ch
			}
		}

		set[string(minS)] = struct{}{}
	}

	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 是所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。

## 专题训练

见下面字符串题单的「**五、最小表示法**」。

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

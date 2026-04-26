> 不考虑返回值，这个做法只需要 $\mathcal{O}(|\Sigma|)$ 空间，其中 $|\Sigma|=5$。

首先，统计 $s$ 包含哪些元音，以及这些元音的出现次数。

然后，把这些元音按照其出现次数从大到小排序。

> 这里的排序是**稳定排序**，这样在出现次数相同时，会自动把位置靠前的元音排在前面。

最后重新填入元音，每填一个元音，就把该元音的出现次数减一。该元音的出现次数为 $0$ 时，切换到下一种元音，作为后续要填入的元音。

[本题视频讲解](https://www.bilibili.com/video/BV1xzZcBZEpe/?t=4m13s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sortVowels(self, s: str) -> str:
        cnt = {}
        vowels = []  # 长度至多为 5
        for ch in s:
            if ch not in "aeiou":
                continue
            if ch not in cnt:
                cnt[ch] = 1
                vowels.append(ch)
            else:
                cnt[ch] += 1

        # 把 aeiou 按照出现次数从大到小排序
        vowels.sort(key=lambda ch: -cnt[ch])

        t = list(s)
        j = 0
        for i, ch in enumerate(t):
            if ch not in "aeiou":
                continue
            t[i] = c = vowels[j]
            cnt[c] -= 1
            if cnt[c] == 0:
                j += 1  # c 消耗完了，切换到下一种元音
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    private static final int[] mp = new int[128];

    static {
        Arrays.fill(mp, -1);
        mp['a'] = 0;
        mp['e'] = 1;
        mp['i'] = 2;
        mp['o'] = 3;
        mp['u'] = 4;
    }

    public String sortVowels(String s) {
        char[] t = s.toCharArray();

        int[] cnt = new int[5];
        List<Character> vowels = new ArrayList<>(); // 长度至多为 5
        for (char ch : t) {
            int x = mp[ch];
            if (x < 0) {
                continue;
            }
            if (cnt[x] == 0) {
                vowels.add(ch);
            }
            cnt[x]++;
        }

        // 把 aeiou 按照出现次数从大到小排序
        vowels.sort((a, b) -> cnt[mp[b]] - cnt[mp[a]]);

        int j = 0;
        for (int i = 0; i < t.length; i++) {
            if (mp[t[i]] < 0) {
                continue;
            }
            t[i] = vowels.get(j);
            if (--cnt[mp[t[i]]] == 0) {
                j++; // 消耗完了，切换到下一种元音
            }
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr array<int, 128> mp = [] {
        array<int, 128> a;
        a.fill(-1);
        a['a'] = 0;
        a['e'] = 1;
        a['i'] = 2;
        a['o'] = 3;
        a['u'] = 4;
        return a;
    }();

public:
    string sortVowels(string s) {
        int cnt[5]{};
        string vowels; // 长度至多为 5
        for (char ch : s) {
            int x = mp[ch];
            if (x < 0) {
                continue;
            }
            if (cnt[x] == 0) {
                vowels += ch;
            }
            cnt[x]++;
        }

        // 把 aeiou 按照出现次数从大到小排序
        ranges::stable_sort(vowels, {}, [&](char ch) { return -cnt[mp[ch]]; });

        int j = 0;
        for (char& ch : s) {
            if (mp[ch] < 0) {
                continue;
            }
            ch = vowels[j];
            if (--cnt[mp[ch]] == 0) {
                j++; // 消耗完了，切换到下一种元音
            }
        }
        return s;
    }
};
```

```go [sol-Go]
var mp = ['z' + 1]int{'a': 1, 'e': 2, 'i': 3, 'o': 4, 'u': 5}

func sortVowels(s string) string {
	cnt := [5]int{}
	vowels := []byte{} // 长度至多为 5
	for _, ch := range s {
		x := mp[ch] - 1
		if x < 0 {
			continue
		}
		if cnt[x] == 0 {
			vowels = append(vowels, byte(ch))
		}
		cnt[x]++
	}

	// 把 aeiou 按照出现次数从大到小排序
	slices.SortStableFunc(vowels, func(a, b byte) int { return cnt[mp[b]-1] - cnt[mp[a]-1] })

	t := []byte(s)
	j := 0
	for i, ch := range t {
		if mp[ch] == 0 {
			continue
		}
		t[i] = vowels[j]
		x := mp[t[i]] - 1
		cnt[x]--
		if cnt[x] == 0 {
			j++ // 消耗完了，切换到下一种元音
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=5$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$ 或者 $\mathcal{O}(|\Sigma|)$，C++ 可以原地修改。

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

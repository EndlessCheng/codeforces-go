题目保证 $p$ 中**恰好**有 $2$ 个星号，这把 $p$ 拆分成了三个子串 $p_1,p_2,p_3$。

**核心思路**：枚举 $p_2$ 在 $s$ 中的位置，同时计算（维护）左边最近的 $p_1$ 的位置，和右边最近的 $p_3$ 的位置，用对应的子串长度更新答案的最小值。

计算 $p_i$ 在 $s$ 中的位置，可以用字符串匹配算法，如 KMP、Z 函数、字符串哈希等，下面用的是 **KMP**，[原理讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

计算左右最近的 $p_1$ 和 $p_3$ 的位置可以用**三指针**维护。

**避免分类讨论的技巧**：如果 $p_i$ 是空的，那么我们认为 $s$ 的所有位置都能匹配空串（包括 $|s|$）。

[本题视频讲解](https://www.bilibili.com/video/BV1BRAGevERN/?t=29m4s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def shortestMatchingSubstring(self, s: str, p: str) -> int:
        p1, p2, p3 = p.split('*')

        # 三段各自在 s 中的所有匹配位置
        pos1 = self.kmp_search(s, p1)
        pos2 = self.kmp_search(s, p2)
        pos3 = self.kmp_search(s, p3)

        ans = inf
        i = k = 0
        # 枚举中间（第二段），维护最近的左右（第一段和第三段）
        for j in pos2:
            # 右边找离 j 最近的子串（但不能重叠）
            while k < len(pos3) and pos3[k] < j + len(p2):
                k += 1
            if k == len(pos3):  # 右边没有
                break
            # 左边找离 j 最近的子串（但不能重叠）
            while i < len(pos1) and pos1[i] <= j - len(p1):
                i += 1
            # 循环结束后，pos1[i-1] 是左边离 j 最近的子串下标（首字母在 s 中的下标）
            if i > 0:
                ans = min(ans, pos3[k] + len(p3) - pos1[i - 1])
        return -1 if ans == inf else ans

    # 计算字符串 p 的 pi 数组
    def calc_pi(self, p: str) -> List[int]:
        pi = [0] * len(p)
        cnt = 0
        for i in range(1, len(p)):
            v = p[i]
            while cnt > 0 and p[cnt] != v:
                cnt = pi[cnt - 1]
            if p[cnt] == v:
                cnt += 1
            pi[i] = cnt
        return pi

    # 在文本串 s 中查找模式串 p，返回所有成功匹配的位置（p[0] 在 s 中的下标）
    def kmp_search(self, s: str, p: str) -> List[int]:
        if not p:
            # s 的所有位置都能匹配空串，包括 len(s)
            return list(range(len(s) + 1))

        pi = self.calc_pi(p)
        pos = []
        cnt = 0
        for i, v in enumerate(s):
            while cnt > 0 and p[cnt] != v:
                cnt = pi[cnt - 1]
            if p[cnt] == v:
                cnt += 1
            if cnt == len(p):
                pos.append(i - len(p) + 1)
                cnt = pi[cnt - 1]
        return pos
```

```java [sol-Java]
class Solution {
    public int shortestMatchingSubstring(String S, String p) {
        char[] s = S.toCharArray();
        String[] sp = p.split("\\*", -1);
        char[] p1 = sp[0].toCharArray();
        char[] p2 = sp[1].toCharArray();
        char[] p3 = sp[2].toCharArray();

        // 三段各自在 s 中的所有匹配位置
        List<Integer> pos1 = kmpSearch(s, p1);
        List<Integer> pos2 = kmpSearch(s, p2);
        List<Integer> pos3 = kmpSearch(s, p3);

        int ans = Integer.MAX_VALUE;
        int i = 0;
        int k = 0;
        // 枚举中间（第二段），维护最近的左右（第一段和第三段）
        for (int j : pos2) {
            // 右边找离 j 最近的子串（但不能重叠）
            while (k < pos3.size() && pos3.get(k) < j + p2.length) {
                k++;
            }
            if (k == pos3.size()) { // 右边没有
                break;
            }
            // 左边找离 j 最近的子串（但不能重叠）
            while (i < pos1.size() && pos1.get(i) <= j - p1.length) {
                i++;
            }
            // 循环结束后，pos1.get(i-1) 是左边离 j 最近的子串下标（首字母在 s 中的下标）
            if (i > 0) {
                ans = Math.min(ans, pos3.get(k) + p3.length - pos1.get(i - 1));
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }

    // 计算字符串 p 的 pi 数组
    private int[] calcPi(char[] p) {
        int[] pi = new int[p.length];
        int match = 0;
        for (int i = 1; i < p.length; i++) {
            char v = p[i];
            while (match > 0 && p[match] != v) {
                match = pi[match - 1];
            }
            if (p[match] == v) {
                match++;
            }
            pi[i] = match;
        }
        return pi;
    }

    // 在文本串 s 中查找模式串 p，返回所有成功匹配的位置（p[0] 在 s 中的下标）
    private List<Integer> kmpSearch(char[] s, char[] p) {
        if (p.length == 0) {
            // s 的所有位置都能匹配空串，包括 s.length
            List<Integer> pos = new ArrayList<>(s.length + 1);
            for (int i = 0; i <= s.length; i++) {
                pos.add(i);
            }
            return pos;
        }

        int[] pi = calcPi(p);
        List<Integer> pos = new ArrayList<>();
        int match = 0;
        for (int i = 0; i < s.length; i++) {
            char v = s[i];
            while (match > 0 && p[match] != v) {
                match = pi[match - 1];
            }
            if (p[match] == v) {
                match++;
            }
            if (match == p.length) {
                pos.add(i - p.length + 1);
                match = pi[match - 1];
            }
        }
        return pos;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 计算字符串 p 的 pi 数组
    vector<int> calcPi(string& p) {
        vector<int> pi(p.size());
        int match = 0;
        for (int i = 1; i < p.size(); i++) {
            char v = p[i];
            while (match > 0 && p[match] != v) {
                match = pi[match - 1];
            }
            if (p[match] == v) {
                match++;
            }
            pi[i] = match;
        }
        return pi;
    }

    // 在文本串 s 中查找模式串 p，返回所有成功匹配的位置（p[0] 在 s 中的下标）
    vector<int> kmp_search(string& s, string p) {
        if (p.empty()) {
            // s 的所有位置都能匹配空串，包括 s.size()
            vector<int> pos(s.size() + 1);
            iota(pos.begin(), pos.end(), 0);
            return pos;
        }

        vector<int> pi = calcPi(p);
        vector<int> pos;
        int match = 0;
        for (int i = 0; i < s.size(); i++) {
            char v = s[i];
            while (match > 0 && p[match] != v) {
                match = pi[match - 1];
            }
            if (p[match] == v) {
                match++;
            }
            if (match == p.size()) {
                pos.push_back(i - p.size() + 1);
                match = pi[match - 1];
            }
        }
        return pos;
    }

public:
    int shortestMatchingSubstring(string s, string p) {
        int star1 = p.find('*');
        int star2 = p.rfind('*');

        // 三段各自在 s 中的所有匹配位置
        vector<int> pos1 = kmp_search(s, p.substr(0, star1));
        vector<int> pos2 = kmp_search(s, p.substr(star1 + 1, star2 - star1 - 1));
        vector<int> pos3 = kmp_search(s, p.substr(star2 + 1));

        // 每一段的长度
        int len1 = star1;
        int len2 = star2 - star1 - 1;
        int len3 = p.size() - star2 - 1;

        int ans = INT_MAX;
        int i = 0, k = 0;
        // 枚举中间（第二段），维护最近的左右（第一段和第三段）
        for (int j : pos2) {
            // 右边找离 j 最近的子串（但不能重叠）
            while (k < pos3.size() && pos3[k] < j + len2) {
                k++;
            }
            if (k == pos3.size()) { // 右边没有
                break;
            }
            // 左边找离 j 最近的子串（但不能重叠）
            while (i < pos1.size() && pos1[i] <= j - len1) {
                i++;
            }
            // 循环结束后，pos1[i-1] 是左边离 j 最近的子串下标（首字母在 s 中的下标）
            if (i > 0) {
                ans = min(ans, pos3[k] + len3 - pos1[i - 1]);
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
// 计算字符串 p 的 pi 数组
func calcPi(p string) []int {
	pi := make([]int, len(p))
	match := 0
	for i := 1; i < len(pi); i++ {
		v := p[i]
		for match > 0 && p[match] != v {
			match = pi[match-1]
		}
		if p[match] == v {
			match++
		}
		pi[i] = match
	}
	return pi
}

// 在文本串 s 中查找模式串 p，返回所有成功匹配的位置（p[0] 在 s 中的下标）
func kmpSearch(s, p string) (pos []int) {
	if p == "" {
		// s 的所有位置都能匹配空串，包括 len(s)
		pos = make([]int, len(s)+1)
		for i := range pos {
			pos[i] = i
		}
		return
	}
	pi := calcPi(p)
	match := 0
	for i := range s {
		v := s[i]
		for match > 0 && p[match] != v {
			match = pi[match-1]
		}
		if p[match] == v {
			match++
		}
		if match == len(p) {
			pos = append(pos, i-len(p)+1)
			match = pi[match-1]
		}
	}
	return
}

func shortestMatchingSubstring(s, p string) int {
	sp := strings.Split(p, "*")
	p1, p2, p3 := sp[0], sp[1], sp[2]

	// 三段各自在 s 中的所有匹配位置
	pos1 := kmpSearch(s, p1)
	pos2 := kmpSearch(s, p2)
	pos3 := kmpSearch(s, p3)

	ans := math.MaxInt
	i, k := 0, 0
	// 枚举中间（第二段），维护最近的左右（第一段和第三段）
	for _, j := range pos2 {
		// 右边找离 j 最近的子串（但不能重叠）
		for k < len(pos3) && pos3[k] < j+len(p2) {
			k++
		}
		if k == len(pos3) { // 右边没有
			break
		}
		// 左边找离 j 最近的子串（但不能重叠）
		for i < len(pos1) && pos1[i] <= j-len(p1) {
			i++
		}
		// 循环结束后，posL[i-1] 是左边离 j 最近的子串下标（首字母在 s 中的下标）
		if i > 0 {
			ans = min(ans, pos3[k]+len(p3)-pos1[i-1])
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $p$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

更多相似题目，见下面字符串题单中的「**一、KMP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. 【本题相关】[字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

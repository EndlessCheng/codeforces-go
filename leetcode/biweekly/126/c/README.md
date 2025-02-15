## 分析

对于每个字母 $i$，统计其出现次数 $\textit{freq}[i]$。

对于 $\texttt{cost}$ 函数的结果 $\textit{res}$ 来说，字母 $i$ 对 $\texttt{res}$ 的贡献就是 $1+2+\cdots + (\textit{freq}[i]-1) = \dfrac{\textit{freq}[i](\textit{freq}[i]-1)}{2}$。

由于问号的出现次数是一个定值，由基本不等式可知，要使 $\textit{res}$ 尽量小，这 $26$ 种字母的出现次数应当尽量接近。

## 方法一：最小堆

1. 统计字母出现次数 $\textit{freq}$，和字母组成 pair 加到一个最小堆中。
2. 设问号出现了 $q$ 次。循环 $q$ 次，每次取出堆顶字母（它是目前出现次数最小的）加入一个列表 $t$ 中，然后把该字母的出现次数加一，重新入堆。
3. 把 $t$ 从小到大排序，因为题目要求字典序最小。
4. 遍历 $s$ 中的问号，按顺序填入 $t$ 中的字母。

[视频讲解](https://www.bilibili.com/video/BV11x421r7q5/) 第三题。

```py [sol-Python3]
class Solution:
    def minimizeStringValue(self, s: str) -> str:
        freq = Counter(s)
        h = [(freq[c], c) for c in ascii_lowercase]
        heapify(h)

        t = []
        for _ in range(s.count('?')):
            f, c = h[0]
            t.append(c)
            heapreplace(h, (f + 1, c))  # 出现次数加一
        t.sort()  # 排序，因为要求字典序最小

        s = list(s)
        j = 0
        for i in range(len(s)):
            if s[i] == '?':
                s[i] = t[j]  # 填入字母
                j += 1
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String minimizeStringValue(String S) {
        char[] s = S.toCharArray();
        int[] freq = new int[26];
        int q = 0;
        for (char c : s) {
            if (c != '?') {
                freq[c - 'a']++;
            } else {
                q++;
            }
        }
        PriorityQueue<Pair<Integer, Character>> pq = new PriorityQueue<>(26, (a, b) -> {
            int c = a.getKey().compareTo(b.getKey());
            return c != 0 ? c : a.getValue().compareTo(b.getValue());
        });
        for (char c = 'a'; c <= 'z'; c++) {
            pq.add(new Pair<>(freq[c - 'a'], c));
        }

        char[] t = new char[q];
        for (int i = 0; i < q; i++) {
            Pair<Integer, Character> p = pq.poll();
            char c = p.getValue();
            t[i] = c;
            pq.add(new Pair<>(p.getKey() + 1, c)); // 出现次数加一
        }
        Arrays.sort(t); // 排序，因为要求字典序最小

        for (int i = 0, j = 0; i < s.length; i++) {
            if (s[i] == '?') {
                s[i] = t[j++]; // 填入字母
            }
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string minimizeStringValue(string s) {
        int freq[26]{};
        for (char c : s) {
            if (c != '?') {
                freq[c - 'a']++;
            }
        }
        priority_queue<pair<int, char>, vector<pair<int, char>>, greater<>> pq;
        for (int i = 0; i < 26; i++) {
            pq.emplace(freq[i], 'a' + i);
        }

        int q = ranges::count(s, '?');
        string t(q, 0);
        for (int i = 0; i < q; i++) {
            auto [f, c] = pq.top();
            pq.pop();
            t[i] = c;
            pq.emplace(f + 1, c); // 出现次数加一
        }
        ranges::sort(t); // 排序，因为要求字典序最小

        for (int i = 0, j = 0; i < s.size(); i++) {
            if (s[i] == '?') {
                s[i] = t[j++]; // 填入字母
            }
        }
        return s;
    }
};
```

```go [sol-Go]
func minimizeStringValue(s string) string {
	h := make(hp, 26)
	for i := byte(0); i < 26; i++ {
		h[i].c = 'a' + i
	}
	for _, b := range s {
		if b != '?' {
			h[b-'a'].f++
		}
	}
	heap.Init(&h)

	t := make([]byte, strings.Count(s, "?"))
	for i := range t {
		t[i] = h[0].c
		h[0].f++ // 出现次数加一
		heap.Fix(&h, 0)
	}
	slices.Sort(t) // 排序，因为要求字典序最小

	ans := []byte(s)
	j := 0
	for i, b := range ans {
		if b == '?' {
			ans[i] = t[j] // 填入字母
			j++
		}
	}
	return string(ans)
}

type pair struct {
	f int
	c byte
}
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.f < b.f || a.f == b.f && a.c < b.c }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。

## 方法二：贪心

下面介绍更快的做法，时间复杂度可以视作线性。

举例说明。为方便描述，假设字母表只有 $\texttt{`a'}$ 到 $\texttt{`d'}$ 四种字母。

1. 统计每种字母的出现次数 $\textit{freq}$，假设为 $[6,8,4,2]$。
2. 复制一份 $f = \textit{freq}$，将 $f$ 从小到大排序，得 $f=[2,4,6,8]$。
3. 设问号的个数 $q=5$。按照方法一的做法，我们实际上要找到一个最小的数 $\textit{limit}$，把小于 $\textit{limit}$ 的出现次数改成 $\textit{limit}$。遍历 $f$ 数组可以计算出，当 $\textit{limit}=5$ 时，我们可以把 $f$ 中的 $2$ 变成 $5$，把 $4$ 变成 $5$，这一共会消耗 $4$ 个问号。多出的 $\textit{extra}=1$ 个问号，给到目前出现次数不超过 $\textit{limit}$ 的最小字母，即 $\texttt{`c'}$。
4. 按照上述做法，创建一个数组 $\textit{target}$，作为替换完所有问号后，每种字母的出现次数。
5. 遍历字符串 $s$，比较 $\textit{freq}$ 和 $\textit{target}$ 来替换字母。如果 $\textit{freq}[j] < \textit{target}[j]$，说明还可以使用字母 $j$，把问号替换成第 $j$ 个字母，然后把 $\textit{freq}[j]$ 增加一。否则，如果 $\textit{freq}[j] = \textit{target}[j]$，说明无法使用字母 $j$，那么增大 $j$，使用后面的字母。

```py [sol-Python3]
class Solution:
    def minimizeStringValue(self, s: str) -> str:
        freq = [0] * 26
        for c in s:
            if c != '?':
                freq[ord(c) - ord('a')] += 1

        f = sorted(freq) + [inf]  # 哨兵
        q = s.count('?')
        for i in count(1):
            need = i * (f[i] - f[i - 1])
            if q <= need:
                limit, extra = f[i - 1] + q // i, q % i
                break
            q -= need

        target = freq.copy()
        for i in range(26):
            if target[i] > limit:
                continue
            target[i] = limit
            if extra:  # 还可以多分配一个
                extra -= 1
                target[i] += 1

        ans = list(s)
        j = 0
        for i, c in enumerate(ans):
            if c != '?':
                continue
            while freq[j] == target[j]:
                j += 1
            freq[j] += 1
            ans[i] = ascii_lowercase[j]
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String minimizeStringValue(String S) {
        char[] s = S.toCharArray();
        int[] freq = new int[27];
        freq[26] = Integer.MAX_VALUE / 26; // 哨兵
        int q = 0;
        for (char c : s) {
            if (c != '?') {
                freq[c - 'a']++;
            } else {
                q++;
            }
        }

        int[] f = freq.clone();
        Arrays.sort(f);

        int limit, extra;
        for (int i = 1; ; i++) {
            int sum = i * (f[i] - f[i - 1]);
            if (q <= sum) {
                limit = f[i - 1] + q / i;
                extra = q % i;
                break;
            }
            q -= sum;
        }

        int[] target = freq.clone();
        for (int j = 0; j < 26; j++) {
            if (freq[j] > limit) {
                continue;
            }
            target[j] = limit;
            if (extra > 0) { // 还可以多分配一个
                extra--;
                target[j]++;
            }
        }

        int j = 0;
        for (int i = 0; i < s.length; i++) {
            if (s[i] != '?') {
                continue;
            }
            while (freq[j] == target[j]) {
                j++;
            }
            freq[j]++;
            s[i] = (char) ('a' + j);
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string minimizeStringValue(string s) {
        array<int, 27> freq;
        freq[26] = INT_MAX / 26; // 哨兵
        for (char c: s) {
            if (c != '?') {
                freq[c - 'a']++;
            }
        }

        auto f = freq;
        ranges::sort(f);

        int limit, extra;
        int q = ranges::count(s, '?');
        for (int i = 1;; i++) {
            int sum = i * (f[i] - f[i - 1]);
            if (q <= sum) {
                limit = f[i - 1] + q / i;
                extra = q % i;
                break;
            }
            q -= sum;
        }

        auto target = freq;
        for (int j = 0; j < 26; j++) {
            if (freq[j] > limit) {
                continue;
            }
            target[j] = limit;
            if (extra > 0) { // 还可以多分配一个
                extra--;
                target[j]++;
            }
        }

        char j = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] != '?') {
                continue;
            }
            while (freq[j] == target[j]) {
                j++;
            }
            freq[j]++;
            s[i] = 'a' + j;
        }
        return s;
    }
};
```

```go [sol-Go]
func minimizeStringValue(s string) string {
	freq := [27]int{26: math.MaxInt / 26} // 哨兵
	for _, c := range s {
		if c != '?' {
			freq[c-'a']++
		}
	}

	f := freq
	slices.Sort(f[:])

	var limit, extra int
	q := strings.Count(s, "?")
	for i := 1; ; i++ {
		sum := i * (f[i] - f[i-1])
		if q <= sum {
			limit, extra = f[i-1]+q/i, q%i
			break
		}
		q -= sum
	}

	target := freq
	for i, c := range freq[:26] {
		if c > limit {
			continue
		}
		target[i] = limit
		if extra > 0 { // 还可以多分配一个
			extra--
			target[i]++
		}
	}

	ans := []byte(s)
	j := byte(0)
	for i, c := range ans {
		if c != '?' {
			continue
		}
		for freq[j] == target[j] {
			j++
		}
		freq[j]++
		ans[i] = 'a' + j
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$，或者 $\mathcal{O}(|\Sigma|)$，如果可以像 C++ 那样原地修改的话。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

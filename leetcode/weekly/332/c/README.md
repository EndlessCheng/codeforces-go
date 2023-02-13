本题 [视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/) 已出炉，欢迎一键三连！

---

把等式 $\textit{val} \oplus \textit{first} = \textit{second}$ 两边同时异或 $\textit{first}$，得到

$$
\textit{val} \oplus \textit{first} \oplus \textit{first} = \textit{second} \oplus \textit{first}
$$

由于 $\textit{first} \oplus \textit{first} = 0$，因此上式化简为

$$
\textit{val} = \textit{second}\oplus \textit{first}
$$

所以问题等价于在 $s$ 中找到值为 $\textit{second}\oplus \textit{first}$ 的数。

由于 $10^9<2^{30}$，我们可以直接预计算所有 $s$ 中长度不超过 $30$ 的数及其对应的 $\textit{left}$ 和 $\textit{right}$，记到一个哈希表中，然后 $O(1)$ 地回答询问。

```py [sol1-Python3]
class Solution:
    def substringXorQueries(self, s: str, queries: List[List[int]]) -> List[List[int]]:
        n, m = len(s), {}
        if (i := s.find('0')) >= 0:
            m[0] = (i, i)  # 这样下面就可以直接跳过 '0' 了，效率更高
        for l, c in enumerate(s):
            if c == '0': continue
            x = 0
            for r in range(l, min(l + 30, n)):
                x = (x << 1) | (ord(s[r]) & 1)
                if x not in m or r - l < m[x][1] - m[x][0]:
                    m[x] = (l, r)

        NOT_FOUND = (-1, -1)
        return [m.get(x ^ y, NOT_FOUND) for x, y in queries]
```

```java [sol1-Java]
class Solution {
    private static final int[] NOT_FOUND = new int[]{-1, -1};

    public int[][] substringXorQueries(String S, int[][] queries) {
        var m = new HashMap<Integer, int[]>();
        int i = S.indexOf('0');
        if (i >= 0) m.put(0, new int[]{i, i}); // 这样下面就可以直接跳过 '0' 了，效率更高
        var s = S.toCharArray();
        for (int l = 0, n = s.length; l < n; ++l) {
            if (s[l] == '0') continue;
            for (int r = l, x = 0; r < Math.min(l + 30, n); ++r) {
                x = x << 1 | (s[r] & 1);
                if (!m.containsKey(x) || r - l < m.get(x)[1] - m.get(x)[0])
                    m.put(x, new int[]{l, r});
            }
        }

        var ans = new int[queries.length][];
        for (i = 0; i < queries.length; i++)
            ans[i] = m.getOrDefault(queries[i][0] ^ queries[i][1], NOT_FOUND);
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<vector<int>> substringXorQueries(string s, vector<vector<int>> &queries) {
        unordered_map<int, pair<int, int>> m;
        if (auto i = s.find('0'); i != string::npos)
            m[0] = {i, i}; // 这样下面就可以直接跳过 '0' 了，效率更高
        for (int l = 0, n = s.length(); l < n; ++l) {
            if (s[l] == '0') continue;
            for (int r = l, x = 0; r < min(l + 30, n); ++r) {
                x = x << 1 | (s[r] & 1);
                auto it = m.find(x);
                if (it == m.end() || r - l < it->second.second - it->second.first)
                    m[x] = {l, r};
            }
        }

        vector<vector<int>> ans;
        for (auto &q : queries) {
            auto it = m.find(q[0] ^ q[1]);
            if (it == m.end()) ans.push_back({-1, -1});
            else ans.push_back({it->second.first, it->second.second});
        }
        return ans;
    }
};
```

```go [sol1-Go]
func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct{ l, r int }
	m := map[int]pair{}
	if i := strings.IndexByte(s, '0'); i >= 0 {
		m[0] = pair{i, i} // 这样下面就可以直接跳过 '0' 了，效率更高
	}
	for l, c := range s {
		if c == '0' {
			continue
		}
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = x<<1 | int(s[r]&1)
			if p, ok := m[x]; !ok || r-l < p.r-p.l {
				m[x] = pair{l, r}
			}
		}
	}

	ans := make([][]int, len(queries))
	notFound := []int{-1, -1} // 避免重复创建
	for i, q := range queries {
		p, ok := m[q[0]^q[1]]
		if !ok {
			ans[i] = notFound
		} else {
			ans[i] = []int{p.l, p.r}
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n\log U + q)$，其中 $n$ 为 $s$ 的长度，$U=max(\textit{queries})$，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n\log U)$。

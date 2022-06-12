#### 提示 1

什么情况下得到的新名字会在 $\textit{ideas}$ 中？

#### 提示 2

按照除去首字母的子串 $\textit{ideas}[i][1:]$ 分组，记录每组的首字母有哪些。

#### 提示 3

不能选两个在同一组的字符串。那么考虑选两个不同组的字符串，这两个字符串需要满足什么要求？

#### 提示 4

设 $\textit{idea}_A$ 的首字母为 $i$，$\textit{idea}_B$ 的首字母为 $j$。那么 $i$ 不能出现在 $\textit{idea}_B$ 的组的首字母中，且 $j$ 不能出现在 $\textit{idea}_A$ 的组的首字母中。

#### 提示 5

尝试枚举 $i$ 和 $j$。我们需要统计什么？

#### 提示 6

定义 $\textit{cnt}[i][j]$ 表示组中首字母不包含 $i$ 但包含 $j$ 的组的个数。枚举每个组，统计 $\textit{cnt}$，同时枚举该组的首字母 $i$ 和不在该组的首字母 $j$，答案即为 $\textit{cnt}[i][j]$ 的累加值。

由于名字需要考虑顺序，最后需要把答案乘 $2$。

```Python [sol1-Python3]
class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        group = defaultdict(int)
        for s in ideas:
            group[s[1:]] |= 1 << (ord(s[0]) - ord('a'))
        ans = 0
        cnt = [[0] * 26 for _ in range(26)]
        for mask in group.values():
            for i in range(26):
                if mask >> i & 1 == 0:
                    for j in range(26):
                        if mask >> j & 1:
                            cnt[i][j] += 1
                else:
                    for j in range(26):
                        if mask >> j & 1 == 0:
                            ans += cnt[i][j]
        return ans * 2
```

```java [sol1-Java]
class Solution {
    public long distinctNames(String[] ideas) {
        var group = new HashMap<String, Integer>();
        for (var s : ideas) {
            var t = s.substring(1);
            group.put(t, group.getOrDefault(t, 0) | 1 << (s.charAt(0) - 'a'));
        }
        var ans = 0L;
        var cnt = new int[26][26];
        for (var mask : group.values())
            for (var i = 0; i < 26; i++)
                if ((mask >> i & 1) == 0) {
                    for (var j = 0; j < 26; j++)
                        if ((mask >> j & 1) > 0) ++cnt[i][j];
                } else {
                    for (var j = 0; j < 26; j++)
                        if ((mask >> j & 1) == 0) ans += cnt[i][j];
                }
        return ans * 2;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long distinctNames(vector<string> &ideas) {
        unordered_map<string, int> group;
        for (auto &s : ideas) {
            auto t = s.substr(1);
            group[t] |= 1 << (s[0] - 'a');
        }
        long ans = 0L;
        int cnt[26][26]; memset(cnt, 0, sizeof(cnt));
        for (auto &[_, mask] : group)
            for (int i = 0; i < 26; i++)
                if ((mask >> i & 1) == 0) {
                    for (int j = 0; j < 26; j++)
                        if (mask >> j & 1) ++cnt[i][j];
                } else {
                    for (int j = 0; j < 26; j++)
                        if ((mask >> j & 1) == 0) ans += cnt[i][j];
                }
        return ans * 2;
    }
};
```

```go [sol1-Go]
func distinctNames(ideas []string) (ans int64) {
	group := map[string]int{}
	for _, s := range ideas {
		group[s[1:]] |= 1 << (s[0] - 'a')
	}
	cnt := [26][26]int{}
	for _, mask := range group {
		for i := 0; i < 26; i++ {
			if mask>>i&1 == 0 {
				for j := 0; j < 26; j++ {
					if mask>>j&1 > 0 {
						cnt[i][j]++
					}
				}
			} else {
				for j := 0; j < 26; j++ {
					if mask>>j&1 == 0 {
						ans += int64(cnt[i][j])
					}
				}
			}
		}
	}
	return ans * 2
}
```


本题 [视频讲解](https://www.bilibili.com/video/BV1aT41157bh) 已出炉，欢迎点赞三连~

---

#### 提示 1

什么情况下得到的新名字不会在 $\textit{ideas}$ 中？

#### 提示 2

按照除去首字母的子串 $\textit{ideas}[i][1:]$ 分组，记录每组的首字母有哪些。

#### 提示 3

我们不能选两个在同一组的字符串。那么考虑选两个分属不同组的字符串，这两个字符串需要满足什么要求？

#### 提示 4

设 $\textit{idea}_A$ 的首字母为 $i$，$\textit{idea}_B$ 的首字母为 $j$。那么 $i$ 不能出现在 $\textit{idea}_B$ 所属组的首字母中，且 $j$ 也不能出现在 $\textit{idea}_A$ 所属组的首字母中。

#### 提示 5

没有头绪？对于这类字符串问题，可以尝试枚举所有小写字母。

尝试枚举 $i$ 和 $j$。我们需要统计什么？

#### 提示 6

定义 $\textit{cnt}[i][j]$ 表示组中首字母不包含 $i$ 但包含 $j$ 的组的个数。枚举每个组，统计 $\textit{cnt}$，同时枚举该组的首字母 $i$ 和不在该组的首字母 $j$，答案即为 $\textit{cnt}[i][j]$ 的累加值。

简单来说就是「有 $i$ 无 $j$」可以和「无 $i$ 有 $j$」的字符串互换。

由于我们是一次遍历所有组，没有考虑两个字符串的顺序，最后需要把答案乘 $2$，表示 A+B 和 B+A 两种字符串的组合。

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
        for (auto &s : ideas)
            group[s.substr(1)] |= 1 << (s[0] - 'a');
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

另一种实现方式是根据首字母分组。

对于 $A$ 组中的字符串 $s$ 和 $B$ 组中的字符串 $t$，只要 $s[1:]$ 和 $t[1:]$ 其中有一个在另一个组中存在，那么这两个字符串就无法互换首字母，否则可以互换。

设 $A$ 组和 $B$ 组交集的大小为 $m$，则这两个组可以组成的合法答案数为

$$
2(len(A)-m)(len(B)-m)
$$

遍历所有组对，累加答案。

注：相比上面的写法，这种写法会让 Python 跑的飞快，但是其他语言并无太大区别。

```Python [sol2-Python3]
class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        group = [set() for _ in range(26)]
        for s in ideas:
            group[ord(s[0]) - ord('a')].add(s[1:])
        ans = 0
        for a, b in combinations(group, 2):
            m = len(a & b)
            ans += (len(a) - m) * (len(b) - m)
        return ans * 2
```

```java [sol2-Java]
class Solution {
    public long distinctNames(String[] ideas) {
        var group = new Set[26];
        for (int i = 0; i < 26; i++) 
            group[i] = new HashSet<String>();
        for (var s : ideas) 
            group[s.charAt(0) - 'a'].add(s.substring(1));
        var ans = 0L;
        for (var i = 0; i < 26; ++i)
            for (var j = 0; j < i; ++j) {
                var m = 0;
                for (var s : group[i])
                    if (group[j].contains(s)) ++m;
                ans += (long) (group[i].size() - m) * (group[j].size() - m);
            }
        return ans * 2;
    }
}
```

```C++ [sol2-C++]
class Solution {
public:
    long long distinctNames(vector<string> &ideas) {
        unordered_set<string> group[26];
        for (auto &s : ideas)
            group[s[0] - 'a'].emplace(s.substr(1));
        long ans = 0L;
        for (int i = 0; i < 26; ++i)
            for (int j = 0; j < i; ++j) {
                int m = 0;
                for (auto &s : group[i])
                    m += group[j].count(s);
                ans += (long) (group[i].size() - m) * (group[j].size() - m);
            }
        return ans * 2;
    }
};
```

```go [sol2-Go]
func distinctNames(ideas []string) (ans int64) {
	set := [26]map[string]bool{}
	for i := range set {
		set[i] = map[string]bool{}
	}
	for _, s := range ideas {
		set[s[0]-'a'][s[1:]] = true
	}
	for i, a := range set {
		for _, b := range set[:i] {
			m := 0
			for s := range a {
				if b[s] {
					m++
				}
			}
			ans += int64(len(a)-m) * int64(len(b)-m)
		}
	}
	return ans * 2
}
```


暴力枚举 $s$ 的所有子串，看与 $\textit{sub}$ 是否匹配。

匹配需要满足字符相同，或者 $\textit{sub}$ 的字符可以通过 $\textit{mappings}$ 转换到 $s$ 的对应字符上。

可以通过将 $\textit{mappings}$ 转换成哈希表或数组来加速判断的过程。

```Python [sol1-Python3]
class Solution:
    def matchReplacement(self, s: str, sub: str, mappings: List[List[str]]) -> bool:
        mp = set((x, y) for x, y in mappings)
        for i in range(len(sub), len(s) + 1):
            if all(x == y or (x, y) in mp for x, y in zip(sub, s[i - len(sub): i])):
                return True
        return False
```

```go [sol1-Go]
func matchReplacement(s, sub string, mappings [][]byte) bool {
	mp := ['z' + 1]['z' + 1]bool{}
	for _, p := range mappings {
		mp[p[0]][p[1]] = true
	}
next:
	for i := len(sub); i <= len(s); i++ {
		for j, c := range s[i-len(sub) : i] {
			if byte(c) != sub[j] && !mp[sub[j]][c] {
				continue next
			}
		}
		return true
	}
	return false
}
```


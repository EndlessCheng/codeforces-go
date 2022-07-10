本题 [视频讲解](https://www.bilibili.com/video/BV1aU4y1q7BA) 已出炉，欢迎点赞三连~

---

首先，无论怎么移动，由于 `L` 和 `R` 无法互相穿过对方，那么去掉 `_` 后的剩余字符应该是相同的，否则返回 `false`。

然后用双指针遍历 $\textit{start}[i]$ 和 $\textit{target}[j]$，分类讨论：

- 如果当前字符为 `L` 且 $i<j$，那么这个 `L` 由于无法向右移动，返回 `false`；
- 如果当前字符为 `R` 且 $i>j$，那么这个 `R` 由于无法向左移动，返回 `false`。

遍历完，若中途没有返回 `false` 就返回 `true`。

```py [sol1-Python3]
class Solution:
    def canChange(self, start: str, target: str) -> bool:
        if start.replace('_', '') != target.replace('_', ''): return False
        j = 0
        for i, c in enumerate(start):
            if c == '_': continue
            while target[j] == '_': j += 1
            if i != j and (c == 'L') != (i > j): return False
            j += 1
        return True
```

```java [sol1-Java]
class Solution {
    public boolean canChange(String start, String target) {
        if (!start.replaceAll("_", "").equals(target.replaceAll("_", ""))) return false;
        for (int i = 0, j = 0; i < start.length(); ++i) {
            if (start.charAt(i) == '_') continue;
            while (target.charAt(j) == '_') ++j;
            if (i != j && (start.charAt(i) == 'L') != (i > j)) return false;
            ++j;
        }
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool canChange(string start, string target) {
        auto s = start, t = target;
        s.erase(remove(s.begin(), s.end(), '_'), s.end());
        t.erase(remove(t.begin(), t.end(), '_'), t.end());
        if (s != t) return false;
        for (int i = 0, j = 0; i < start.length(); ++i) {
            if (start[i] == '_') continue;
            while (target[j] == '_') ++j;
            if (i != j && (start[i] == 'L') != (i > j)) return false;
            ++j;
        }
        return true;
    }
};
```

```go [sol1-Go]
func canChange(start, target string) bool {
	if strings.ReplaceAll(start, "_", "") != strings.ReplaceAll(target, "_", "") {
		return false
	}
	j := 0
	for i, c := range start {
		if c != '_' {
			for target[j] == '_' {
				j++
			}
			if i != j && c == 'L' != (i > j) {
				return false
			}
			j++
		}
	}
	return true
}
```

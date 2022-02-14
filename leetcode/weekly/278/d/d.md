由于字符串不含重复字母且只由小写字母组成，我们可以用一个二进制数来表示字符串，二进制数的第 $i$ 位为 $1$ 则表示第 $i$ 个小写字母出现在 $s$ 中。

枚举 $\textit{words}$ 中的字符串 $s$，并枚举 $s$ 通过添加、删除和替换操作得到的字符串 $t$，如果 $t$ 也在 $\textit{words}$ 中，则说明 $s$ 和 $t$ 可以分到同一组。我们可以用并查集来关联可以分到同一组的字符串。

遍历结束后，并查集中的集合个数就是 $\textit{words}$ 分组后的总组数，最大的集合即为字符串数目最多的组所包含的字符串数目。这可以在并查集合并的同时维护出来。

```go [sol1-Go]
func groupStrings(words []string) (ans []int) {
	// 并查集模板（哈希表写法）
	fa := map[int]int{}
	size := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	groups, maxSize := len(words), 0
	merge := func(x, y int) {
		if _, ok := fa[y]; !ok {
			return
		}
		x, y = find(x), find(y)
		if x == y {
			return
		}
		fa[x] = y
		size[y] += size[x]
		maxSize = max(maxSize, size[y]) // 维护答案
		groups--
	}

	for _, word := range words {
		x := 0
		for _, ch := range word {
			x |= 1 << (ch - 'a') // 计算 word 的二进制表示
		}
		fa[x] = x // 添加至并查集
		size[x]++
		maxSize = max(maxSize, size[x]) // 维护答案
		if size[x] > 1 {
			groups--
		}
	}

	for x := range fa { // 枚举所有字符串（二进制表示）
		for i := 0; i < 26; i++ {
			merge(x, x^1<<i) // 添加或删除字符 i
			if x>>i&1 == 1 {
				for j := 0; j < 26; j++ {
					if x>>j&1 == 0 {
						merge(x, x^1<<i|1<<j) // 替换字符 i 为 j
					}
				}
			}
		}
	}
	return []int{groups, maxSize}
}

func max(a, b int) int { if b > a { return b }; return a }
```

```C++ [sol1-C++]
class Solution {
    // 并查集模板（哈希表写法）
    unordered_map<int, int> fa, size;
    int groups, maxSize = 0;

    int find(int x) {
        return fa[x] != x ? fa[x] = find(fa[x]) : x;
    }

    void merge(int x, int y) {
        if (!fa.count(y)) return;
        x = find(x);
        y = find(y);
        if (x == y) return;
        fa[x] = y;
        size[y] += size[x];
        maxSize = max(maxSize, size[y]); // 维护答案
        --groups;
    }

public:
    vector<int> groupStrings(vector<string> &words) {
        groups = words.size();
        for (auto &word : words) {
            int x = 0;
            for (char ch: word) {
                x |= 1 << (ch - 'a'); // 计算 word 的二进制表示
            }
            fa[x] = x;  // 添加至并查集
            ++size[x];
            maxSize = max(maxSize, size[x]); // 维护答案
            if (size[x] > 1) --groups;
        }

        for (auto &[x, _]: fa) { // 枚举所有字符串（二进制表示）
            for (int i = 0; i < 26; ++i) {
                merge(x, x ^ (1 << i)); // 添加或删除字符 i
                if ((x >> i) & 1) {
                    for (int j = 0; j < 26; ++j) {
                        if (((x >> j) & 1) == 0) {
                            merge(x, x ^ (1 << i) | (1 << j)); // 替换字符 i 为 j
                        }
                    }
                }
            }
        }
        return {groups, maxSize};
    }
};
```

```Python [sol1-Python3]
class Solution:
    def groupStrings(self, words: List[str]) -> List[int]:
        # 并查集模板（哈希表写法）
        fa, size = {}, defaultdict(int)
        groups, max_size = len(words), 0
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]
        def merge(x: int, y: int):
            nonlocal groups, max_size
            if y not in fa:
                return
            x, y = find(x), find(y)
            if x == y:
                return
            fa[x] = y
            size[y] += size[x]
            max_size = max(max_size, size[y])  # 维护答案
            groups -= 1

        for word in words:
            x = 0
            for ch in word:
                x |= 1 << (ord(ch) - ord('a'))  # 计算 word 的二进制表示
            fa[x] = x  # 添加至并查集
            size[x] += 1
            max_size = max(max_size, size[x])  # 维护答案
            if size[x] > 1:
                groups -= 1

        for x in fa:  # 枚举所有字符串（二进制表示）
            for i in range(26):
                merge(x, x ^ (1 << i))  # 添加或删除字符 i
                if (x >> i) & 1:
                    for j in range(26):
                        if ((x >> j) & 1) == 0:
                            merge(x, x ^ (1 << i) | (1 << j))  # 替换字符 i 为 j
        return [groups, max_size]
```

```java [sol1-Java]
class Solution {
    // 并查集模板（哈希表写法）
    HashMap<Integer, Integer> fa = new HashMap<>(), size = new HashMap<>();
    int groups, maxSize;

    int find(int x) {
        if (fa.get(x) != x) {
            fa.put(x, find(fa.get(x)));
        }
        return fa.get(x);
    }

    void merge(int x, int y) {
        if (!fa.containsKey(y)) return;
        x = find(x);
        y = find(y);
        if (x == y) return;
        fa.put(x, y);
        size.put(y, size.get(y) + size.get(x));
        maxSize = Math.max(maxSize, size.get(y)); // 维护答案
        --groups;
    }

    public int[] groupStrings(String[] words) {
        groups = words.length;
        for (var word : words) {
            var x = 0;
            for (var i = 0; i < word.length(); i++) {
                x |= 1 << (word.charAt(i) - 'a'); // 计算 word 的二进制表示
            }
            fa.put(x, x); // 添加至并查集
            size.put(x, size.getOrDefault(x, 0) + 1);
            maxSize = Math.max(maxSize, size.get(x)); // 维护答案
            if (size.get(x) > 1) --groups;
        }

        fa.forEach((x, fx) -> {
            for (var i = 0; i < 26; i++) {
                merge(x, x ^ (1 << i)); // 添加或删除字符 i
                if (((x >> i) & 1) == 1) {
                    for (var j = 0; j < 26; ++j) {
                        if (((x >> j) & 1) == 0) {
                            merge(x, x ^ (1 << i) | (1 << j)); // 替换字符 i 为 j
                        }
                    }
                }
            }
        });
        return new int[]{groups, maxSize};
    }
}
```

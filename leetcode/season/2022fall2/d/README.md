[视频讲解](https://www.bilibili.com/video/BV1rT411P7NA) 已出炉，欢迎点赞三连，在评论区分享你对这场力扣杯的看法~

记录关键思路，详细的说明见视频讲解。

1. 用位运算表示字母选择情况，由于一个字母可以选多个，因此要对二进制「分区」，每个区域表示对应字母的个数。
2. 写一个记忆化搜索，$f(i,\textit{mask})$ 表示从第 $i$ 个单词开始选，已经选择的单词为 $\textit{mask}$ 时，后续消耗代币的最小值。枚举 $\textit{words}[i]$ 的所有合法选择方案转移到 $f(i+1,\textit{mask}')$。
3. 因此需要预处理每个 $\textit{words}[i]$ 的每种选择字母的方案所消耗的代币的最小值，由于字符串很短，直接写个爆搜即可。

```py [sol1-Python3]
# (字母在二进制上的起始位置, 这个字母能选择的上限, 位掩码)
RULES = {
    'e': (0, 4, 7),
    'l': (3, 3, 3),
    'o': (5, 2, 3),
    'h': (7, 1, 1),
    't': (8, 1, 1),
    'c': (9, 1, 1),
    'd': (10, 1, 1),
}
FULL = 2012  # 0b11111011100，每个字母都选到了对应的上限

# 合并两种选择字母的方案
def merge(cur: int, add: int) -> int:
    for pos, limit, m in RULES.values():
        c1 = (cur >> pos) & m
        c2 = (add >> pos) & m
        if c1 + c2 > limit:
            return -1
        cur += c2 << pos
    return cur

class Solution:
    def Leetcode(self, words: List[str]) -> int:
        # 预处理每个单词的每种选择字母的方案所消耗的代币的最小值
        costs = []
        for word in words:
            cost = {}
            def dfs(s: str, mask: int, tot: int) -> None:
                if mask not in cost or tot < cost[mask]:
                    cost[mask] = tot
                for i, c in enumerate(s):  # 枚举选择字母的位置
                    if c not in RULES:
                        continue
                    pos, limit, m = RULES[c]
                    if (mask >> pos) & m < limit:  # 可以选字母 c
                        dfs(s[:i] + s[i + 1:], mask + (1 << pos), tot + i * (len(s) - 1 - i))
            dfs(word, 0, 0)
            costs.append(cost)

        @cache
        def dfs(i: int, mask: int) -> int:
            if i == len(words):
                return 0 if mask == FULL else inf  # inf 表示不合法，没有选完要求的字母
            res = inf
            for add, tot in costs[i].items():
                if tot >= res: continue  # 剪枝
                m = merge(mask, add)
                if m >= 0:
                    res = min(res, tot + dfs(i + 1, m))
            return res
        ans = dfs(0, 0)
        return ans if ans < inf else -1
```

```go [sol1-Go]
const keys = "elohtcd"
const full = 2012 // 0b11111011100，每个字母都选到了对应的上限

// pos：字母在二进制上的起始位置
// limit：这个字母能选择的上限
// mask：位掩码
var rules = ['z' + 1]struct{ pos, limit, mask int }{
	'e': {0, 4, 7},
	'l': {3, 3, 3},
	'o': {5, 2, 3},
	'h': {7, 1, 1},
	't': {8, 1, 1},
	'c': {9, 1, 1},
	'd': {10, 1, 1},
}

// 合并两种选择字母的方案
func merge(cur, add int) int {
	for _, c := range keys {
		r := rules[c]
		c1 := cur >> r.pos & r.mask
		c2 := add >> r.pos & r.mask
		if c1+c2 > r.limit {
			return -1
		}
		cur += c2 << r.pos
	}
	return cur
}

func Leetcode(words []string) int {
	const inf = math.MaxInt32 / 2
	n := len(words)
	// 预处理每个单词的每种选择字母的方案所消耗的代币的最小值
	costs := make([][1 << 11]int, n)
	for i, word := range words {
		for j := range costs[i] {
			costs[i][j] = inf
		}
		var f func(string, int, int)
		f = func(s string, mask, tot int) {
			costs[i][mask] = min(costs[i][mask], tot)
			for j, c := range s { // 枚举选择字母的位置
				r := rules[c]
				if mask>>r.pos&r.mask < r.limit { // 可以选字母 c
					f(s[:j]+s[j+1:], mask+1<<r.pos, tot+j*(len(s)-1-j))
				}
			}
		}
		f(word, 0, 0)
	}

	dp := make([][1 << 11]int, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, mask int) int {
		if i == n {
			if mask == full {
				return 0
			}
			return inf // inf 表示不合法，没有选完要求的字母
		}
		ptr := &dp[i][mask]
		if *ptr != -1 {
			return *ptr
		}
		res := inf
		for add, tot := range costs[i] {
			if tot >= res { // 剪枝
				continue
			}
			m2 := merge(mask, add)
			if m2 >= 0 {
				res = min(res, f(i+1, m2)+tot)
			}
		}
		*ptr = res
		return res
	}
	ans := f(0, 0)
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

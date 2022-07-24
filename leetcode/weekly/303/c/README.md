下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

用三个哈希表模拟。

```py [sol1-Python3]
class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.fr = {}
        self.fc = {}
        self.cs = defaultdict(SortedSet)
        for f, c, r in zip(foods, cuisines, ratings):
            self.fr[f] = r
            self.fc[f] = c
            self.cs[c].add((-r, f))

    def changeRating(self, f: str, newRating: int) -> None:
        s = self.cs[self.fc[f]]
        s.remove((-self.fr[f], f))  # 移除旧数据
        s.add((-newRating, f))  # 添加新数据
        self.fr[f] = newRating

    def highestRated(self, cuisine: str) -> str:
        return self.cs[cuisine][0][1]
```

```cpp [sol1-C++]
class FoodRatings {
    unordered_map<string, int> fr;
    unordered_map<string, string> fc;
    unordered_map<string, set<pair<int, string>>> cs;
public:
    FoodRatings(vector<string> &foods, vector<string> &cuisines, vector<int> &ratings) {
        for (int i = 0; i < foods.size(); ++i) {
            auto &f = foods[i], &c = cuisines[i];
            fr[f] = ratings[i];
            fc[f] = c;
            cs[c].emplace(-ratings[i], f);
        }
    }

    void changeRating(string f, int newRating) {
        auto &s = cs[fc[f]];
        s.erase({-fr[f], f}); // 移除旧数据
        s.emplace(-newRating, f); // 添加新数据
        fr[f] = newRating;
    }

    string highestRated(string cuisine) {
        return cs[cuisine].begin()->second;
    }
};
```

```go [sol1-Go]
type pair struct {
	r int
	f string
}

type FoodRatings struct {
	fr map[string]int
	fc map[string]string
	ct map[string]*redblacktree.Tree
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fr := map[string]int{}
	fc := map[string]string{}
	ct := map[string]*redblacktree.Tree{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fr[f] = r
		fc[f] = c
		if ct[c] == nil {
			ct[c] = redblacktree.NewWith(func(x, y interface{}) int {
				a, b := x.(pair), y.(pair)
				if a.r != b.r {
					return utils.IntComparator(b.r, a.r)
				}
				return utils.StringComparator(a.f, b.f)
			})
		}
		ct[c].Put(pair{r, f}, nil)
	}
	return FoodRatings{fr, fc, ct}
}

func (r FoodRatings) ChangeRating(f string, newRating int) {
	t := r.ct[r.fc[f]]
	t.Remove(pair{r.fr[f], f}) // 移除旧数据
	t.Put(pair{newRating, f}, nil) // 添加新数据
	r.fr[f] = newRating
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.ct[cuisine].Left().Key.(pair).f
}
```

package main

// https://space.bilibili.com/206214
func mostPopularCreator(creators, ids []string, views []int) (ans [][]string) {
	type tuple struct {
		viewSum, maxView int
		id               string
	}
	m, maxViewSum := map[string]tuple{}, 0
	for i, name := range creators {
		id, view := ids[i], views[i]
		t := m[name]
		if t.id == "" {
			t = tuple{view, view, id}
		} else {
			t.viewSum += view
			if view > t.maxView || view == t.maxView && id < t.id {
				t.maxView = view
				t.id = id
			}
		}
		maxViewSum = max(maxViewSum, t.viewSum)
		m[name] = t
	}
	for name, t := range m {
		if t.viewSum == maxViewSum {
			ans = append(ans, []string{name, t.id})
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }

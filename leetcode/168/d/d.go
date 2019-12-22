package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) (ans int) {
	n := len(status)
	haveKey := make([]bool, n)
	haveBox := make([]bool, n)
	seen := make([]bool, n)
	for _, b := range initialBoxes {
		haveBox[b] = true
	}
	for {
		foundNewBox := false
		for b, hav := range haveBox {
			if !hav || seen[b] || status[b] == 0 && !haveKey[b] {
				continue
			}
			ans += candies[b]
			for _, k := range keys[b] {
				haveKey[k] = true
			}
			for _, bb := range containedBoxes[b] {
				haveBox[bb] = true
			}
			seen[b] = true
			foundNewBox = true
		}
		if !foundNewBox {
			break
		}
	}
	return
}

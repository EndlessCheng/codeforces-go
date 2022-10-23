package main

// https://space.bilibili.com/206214
func haveConflict(event1, event2 []string) bool {
	return event1[0] <= event2[1] && event1[1] >= event2[0]
}

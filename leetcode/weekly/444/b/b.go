package main

import "sort"

// https://space.bilibili.com/206214
type packet struct {
	source, destination, timestamp int
}

type Router struct {
	memoryLimit      int
	packetQ          []packet            // packet 队列
	packetSet        map[packet]struct{} // packet 集合
	destToTimestamps map[int][]int       // destination -> [timestamp]
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:      memoryLimit,
		packetSet:        map[packet]struct{}{},
		destToTimestamps: map[int][]int{},
	}
}

func (r *Router) AddPacket(source, destination, timestamp int) bool {
	pkt := packet{source, destination, timestamp}
	if _, ok := r.packetSet[pkt]; ok {
		return false
	}
	r.packetSet[pkt] = struct{}{}
	if len(r.packetQ) == r.memoryLimit { // 太多了
		r.ForwardPacket()
	}
	r.packetQ = append(r.packetQ, pkt) // 入队
	r.destToTimestamps[destination] = append(r.destToTimestamps[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.packetQ) == 0 {
		return nil
	}
	pkt := r.packetQ[0]
	r.packetQ = r.packetQ[1:] // 出队
	r.destToTimestamps[pkt.destination] = r.destToTimestamps[pkt.destination][1:]
	delete(r.packetSet, pkt)
	return []int{pkt.source, pkt.destination, pkt.timestamp}
}

func (r *Router) GetCount(destination, startTime, endTime int) int {
	timestamps := r.destToTimestamps[destination]
	return sort.SearchInts(timestamps, endTime+1) - sort.SearchInts(timestamps, startTime)
}

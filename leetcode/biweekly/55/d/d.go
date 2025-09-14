package main

import (
	"cmp"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

// github.com/EndlessCheng/codeforces-go
type shopMovie struct{ shop, movie int }
type priceShop struct{ price, shop int }
type entry struct{ price, shop, movie int }

type MovieRentingSystem struct {
	shopMovieToPrice         map[shopMovie]int
	unrentedMovieToPriceShop map[int]*redblacktree.Tree[priceShop, struct{}]
	rentedMovies             *redblacktree.Tree[entry, struct{}]
}

func Constructor(_ int, entries [][]int) MovieRentingSystem {
	shopMovieToPrice := map[shopMovie]int{}
	unrentedMovieToPriceShop := map[int]*redblacktree.Tree[priceShop, struct{}]{}

	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		shopMovieToPrice[shopMovie{shop, movie}] = price
		if _, ok := unrentedMovieToPriceShop[movie]; !ok {
			unrentedMovieToPriceShop[movie] = redblacktree.NewWith[priceShop, struct{}](func(a, b priceShop) int {
				return cmp.Or(a.price-b.price, a.shop-b.shop)
			})
		}
		unrentedMovieToPriceShop[movie].Put(priceShop{price, shop}, struct{}{})
	}

	rentedMovies := redblacktree.NewWith[entry, struct{}](func(a, b entry) int {
		return cmp.Or(a.price-b.price, a.shop-b.shop, a.movie-b.movie)
	})
	return MovieRentingSystem{shopMovieToPrice, unrentedMovieToPriceShop, rentedMovies}
}

func (m *MovieRentingSystem) Search(movie int) (ans []int) {
	// unrentedMovieToPriceShop[movie] 前 5 个
	t := m.unrentedMovieToPriceShop[movie]
	if t == nil {
		return
	}
	for it := t.Iterator(); len(ans) < 5 && it.Next(); {
		ans = append(ans, it.Key().shop)
	}
	return
}

func (m *MovieRentingSystem) Rent(shop, movie int) {
	price := m.shopMovieToPrice[shopMovie{shop, movie}]
	// 从 unrentedMovieToPriceShop 移到 rentedMovies
	m.unrentedMovieToPriceShop[movie].Remove(priceShop{price, shop})
	m.rentedMovies.Put(entry{price, shop, movie}, struct{}{})
}

func (m *MovieRentingSystem) Drop(shop, movie int) {
	price := m.shopMovieToPrice[shopMovie{shop, movie}]
	// 从 rentedMovies 移到 unrentedMovieToPriceShop
	m.rentedMovies.Remove(entry{price, shop, movie})
	m.unrentedMovieToPriceShop[movie].Put(priceShop{price, shop}, struct{}{})
}

func (m *MovieRentingSystem) Report() (ans [][]int) {
	// rentedMovies 前 5 个
	for it := m.rentedMovies.Iterator(); len(ans) < 5 && it.Next(); {
		ans = append(ans, []int{it.Key().shop, it.Key().movie})
	}
	return
}

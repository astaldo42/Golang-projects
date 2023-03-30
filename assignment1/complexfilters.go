package main

import "sort"

type byPriceDesc []Item

func (a byPriceDesc) Len() int {
	return len(a)
}

func (a byPriceDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byPriceDesc) Less(i, j int) bool {
	return a[i].Price > a[j].Price
}

type byRatingDesc []Item

func (a byRatingDesc) Len() int {
	return len(a)
}

func (a byRatingDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byRatingDesc) Less(i, j int) bool {
	return a[i].Rating > a[j].Rating
}

func FilterItemsByPrice(price float64) []Item {
	var result []Item
	for _, item := range Items {
		if item.Price <= price {
			result = append(result, item)
		}
	}
	sort.Sort(byPriceDesc(result))
	return result
}

func FilterItemsByRating(rating float64) []Item {
	var result []Item
	for _, item := range Items {
		if item.Rating >= rating {
			result = append(result, item)
		}
	}
	sort.Sort(byRatingDesc(result))
	return result
}

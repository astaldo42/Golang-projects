package filters

import (
	"github.com/Jateq/oop2/base"
	"sort"
)

type byPriceDesc []base.Item

func (a byPriceDesc) Len() int {
	return len(a)
}

func (a byPriceDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byPriceDesc) Less(i, j int) bool {
	return a[i].Price > a[j].Price
}

type byRatingDesc []base.Item

func (a byRatingDesc) Len() int {
	return len(a)
}

func (a byRatingDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byRatingDesc) Less(i, j int) bool {
	return a[i].Rating > a[j].Rating
}

func FilterItemsByPrice() []base.Item {
	var price float64 = 1500
	var result []base.Item
	for _, item := range base.Items {
		if item.Price <= price {
			result = append(result, item)
		}
	}
	sort.Sort(byPriceDesc(result))
	return result
}

func FilterItemsByRating() []base.Item {
	var rating float64 = 0
	var result []base.Item
	for _, item := range base.Items {
		if item.Rating >= rating {
			result = append(result, item)
		}
	}
	sort.Sort(byRatingDesc(result))
	return result
}

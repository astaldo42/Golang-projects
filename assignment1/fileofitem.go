package main

import "strconv"

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

var Items []Item

func Store(market []Item) string {
	var result string
	for _, item := range market {
		result += "Name: " + item.Name + "\tPrice: " + strconv.FormatFloat(item.Price, 'f', 2, 64) + "\tRating: " + strconv.FormatFloat(item.Rating, 'f', 2, 64) + "\n"
	}
	return result
}

func SearchItemsByName(name string) []Item {
	var result []Item
	for _, item := range Items {
		if item.Name == name {
			result = append(result, item)
		}
	}
	return result
}

//func FilterItemsByPrice(price float64) []Item {
//	var result []Item
//	for _, item := range Items {
//		if item.Price <= price {
//			result = append(result, item)
//		}
//	}
//	return result
//}

//func FilterItemsByRating(rating float64) []Item {
//	var result []Item
//	for _, item := range Items {
//		if item.Rating >= rating {
//			result = append(result, item)
//		}
//	}
//	return result
//}

func GiveRating(name string, rating float64) {
	for i := range Items {
		if Items[i].Name == name {
			Items[i].Rating = rating
			break
		}
	}
}

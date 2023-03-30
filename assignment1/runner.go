package main

import (
	"fmt"
)

func main() {

	fmt.Println("\033[1;33mWelcome to Jateq's market\033[0m")
	LoadDataFromJSON("users.json", &Users)
	LoadDataFromJSON("items.json", &Items)
	fmt.Println("Do you have an account? (y or n)")
	var input, username, password, corrector string
	for true {
		fmt.Scanln(&input)
		if input == "y" {
			fmt.Println("\033[1;33mPlease log in\033[0m")
			fmt.Scanln(&username, &password)
			if Authorization(username, password) {
				fmt.Print("\033[1;32m")
				fmt.Println("Authorization succeeded for", username)
				fmt.Print("\033[0m")
				break
			} else {
				for !Authorization(username, password) {
					fmt.Println("\033[1;31mAuthorization failed, try again\033[0m")
					fmt.Scanln(&username, &password)
				}
				fmt.Print("\033[1;32m")
				fmt.Println("Authorization succeeded for", username)
				fmt.Print("\033[0m")
				break

			}
		} else if input == "n" {
			fmt.Println("Please register (username, password)")
			fmt.Scanln(&username, &password)
			fmt.Println("Please retype your password again")
			fmt.Scanln(&corrector)
			for password != corrector {
				fmt.Println("They are not match, please make sure everything is ok")
				fmt.Scanln(&corrector)
			}
			Registration(username, password)
			fmt.Println("You are welcome to Jateq's Store")
			break
		} else {
			fmt.Println("\033[1;31mInvalid input. Please enter either 'y' or 'n'\033[0m")
			continue
		}
	}

	fmt.Println("\n\033[1;33mOffers for today:\033[0m")
	fmt.Println(Store(Items))
	//fmt.Println("-----------------------------")

	fmt.Println("\033[1;33mSearch in market:\033[0m")
	var name string
	fmt.Scanln(&name)
	items := SearchItemsByName(name)
	if len(items) == 0 {
		fmt.Println("\033[1;31mSorry", name, "are out of stock\033[0m\n")
	} else {
		for _, item := range items {
			fmt.Println("Name:", item.Name)
			fmt.Println("Price:", item.Price)
			fmt.Println("Rating:", item.Rating)
			fmt.Println("-----------------------------\n")
		}
		var answer string
		var rate float64
		fmt.Println("\033[1;33mDo you want to buy it? (y or n)\033[0m")
		fmt.Scanln(&answer)
		if answer == "y" {
			fmt.Println("\033[1;33mDid you liked it? Give us your rating (1 - 5)\033[0m")
			fmt.Scanln(&rate)
			GiveRating(name, rate)
			fmt.Println("\033[1;32mThanks for your feedback\033[0m")
		} else {
			fmt.Println("\033[1;33mOkay\033[0m")
		}
	}
	var input3, specif string
	var further float64
	for true {
		fmt.Println("\033[1;33mIs there something else that I can help you?\033[0m")
		fmt.Println("\033[1;34mFilter\033[0m    \033[1;32mSearch\033[0m     \033[1;31mExit\033[0m")
		fmt.Scanln(&input3)
		if input3 == "Filter" || input3 == "filter" {
			for true {
				fmt.Println("\033[1;34mFilter by Price or Rate\033[0m")
				fmt.Scanln(&specif)
				if specif == "Price" || specif == "price" {
					fmt.Println("\033[1;34mWhat is your max affordable price\033[0m")
					fmt.Scanln(&further)

					itemsByPrice := FilterItemsByPrice(further)
					fmt.Println("\033[1;34mItems with price less than", further, "\033[0m\n")
					fmt.Println(Store(itemsByPrice))
					break

				} else if specif == "Rate" || specif == "rate" {
					fmt.Println("\033[1;34mWhat is the lower bound for rating\033[0m")
					fmt.Scanln(&further)

					itemsByRating := FilterItemsByRating(further)
					fmt.Println("\033[1;34mItems with rating more than", further, "\033[0m\n")
					fmt.Println(Store(itemsByRating))
					break
				} else {
					fmt.Println("\033[1;31mPlease write Rate or Price only\033[0m")
					continue
				}
			}
		} else if input3 == "Search" || input3 == "search" {
			var name1 string
			fmt.Println("\033[1;32mWhat you want to find?\033[0m")
			fmt.Scanln(&name1)
			items2 := SearchItemsByName(name1)
			if len(items2) == 0 {
				fmt.Println("\033[1;31mSorry", name1, "are out of stock\033[0m\n")
			} else {
				for _, item := range items2 {
					fmt.Println("\n\033[1;32mName:", item.Name)
					fmt.Println("Price:", item.Price)
					fmt.Println("Rating:", item.Rating)
					fmt.Println("-----------------------------\033[0m\n")
					var answer1 string
					var rate1 float64
					fmt.Println("\033[1;33mDo you want to buy it? (y or n)\033[0m")
					fmt.Scanln(&answer1)
					if answer1 == "y" {
						fmt.Println("\033[1;33mDid you liked it? Give us your rating (1 - 5)\033[0m")
						fmt.Scanln(&rate1)
						GiveRating(name, rate1)
						fmt.Println("\033[1;32mThanks for your feedback\033[0m\n")
					} else {
						fmt.Println("\033[1;33mOkay\033[0m\n")
					}
				}
			}
		} else if input3 == "Exit" || input3 == "exit" {
			fmt.Println("\033[1;42m", username, "\033[0m\033[1;33m hope to see you here soon\033[0m ")
			break
		} else {
			fmt.Println("\n\033[1;31mType only this options\033[0m")
			continue
		}
	}
	SaveDataToJSON("users.json", &Users)
	SaveDataToJSON("items.json", &Items)
}

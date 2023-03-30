package database

type Books struct {
	ID          uint
	Title       string
	Description string
	Cost        uint
}

func MyBooks() []Books {
	books := []Books{
		{Title: "Cruel Age", Description: "A book written by Isai K.K about life of Chengis Khan", Cost: 26},
		{Title: "Crime and Punishment", Description: "One of the most gripping crime stories of all time written by Fyodor D.", Cost: 15},
		{Title: "Dune", Description: "Dune, Frank Herbert’s epic science-fiction masterpiece set in the far future amidst a sprawling feudal interstellar society", Cost: 20},
		{Title: "War and Peace", Description: "Epic in scale, War and Peace delineates in graphic detail events leading up to Napoleon’s invasion of Russia, written by Leo Tolstoy", Cost: 23},
		{Title: "Meditations", Description: "Meditations is a series of personal writings by Marcus Aurelius", Cost: 25},
		{Title: "Moby-Dick", Description: "Every American writer since 1851 has been chasing the same whale: to somehow write a novel as epic and influential as Melville’s.", Cost: 27},
		{Title: "Frankenstein", Description: "Written when Mary Shelley was just 18 years old, but don’t let that depress you. Frankenstein is a Gothic masterpiece with entertaining set pieces aplenty.", Cost: 14},
		{Title: "The lord of the Rings", Description: " J. R. R. Tolkien’s incredible trilogy of otherworldliness brought a world of hobbits, dwarves, elves and orcs to life in a way never read before. ", Cost: 50},
		{Title: "Dracula", Description: "A Gothic tale of fear and love by Bram Stocker. Would one desire immortality at the cost of one’s morality and soul?", Cost: 40},
		{Title: "Pride and Prejudice", Description: "It is a truth universally acknowledged that when most people think of Jane Austen they think of this charming and humorous story of love", Cost: 30},
	}
	return books
}

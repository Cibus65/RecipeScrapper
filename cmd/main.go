package main

import (
	"fmt"
	"scrapper/app"
)

func main() {
	var pages int
	fmt.Println("Сколько страниц надо спарсить?")
	fmt.Scan(&pages)
	app.ScrapeRecipe("https://eda.ru/recepty?page=", "https://eda.ru", pages)
	//repo.UpdateDB(domain.Recipe{Name: "d"})
}

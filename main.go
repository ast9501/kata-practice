package main

import (
	"bufio"
	"fmt"
	"os"
	. "potter/internal"
	"strconv"
)

func main() {
	fmt.Println("Enter list of book(s): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	myCalculator := BasicCalculator{
		Discount: []float32{0, 0.95, 0.9, 0.8, 0.75},
	}
	var list []int
	for _, v := range input {
		if string(v) == "\n" {
			continue
		}
		s, _ := strconv.Atoi(string(v))
		list = append(list, s)
	}
	shoppingList := ShoppingList{
		Books:       list,
		CostPerBook: 8.0,
		TotalCost:   0,
	}
	shoppingList = myCalculator.CountCost(shoppingList)
	fmt.Println(shoppingList.TotalCost)
}

package internal

import (
	"sort"
)

type Calculator interface {
	CostCount(ShoppingList) ShoppingList
	SetDiscount() bool
	GetDiscount() []int
}

type BasicCalculator struct {
	Discount []float32
}

type ShoppingList struct {
	Books       []int
	CostPerBook float32
	TotalCost   float32
}

func (c *BasicCalculator) CountCost(s ShoppingList) (r ShoppingList) {

	sort.Ints(s.Books)

	r.TotalCost = 0
	group := make([][]int, 2)
	modifiedGroup := group
	for _, v := range s.Books {
		for k, s := range group {
			if !Contain(s, v) {
				if len(s) <= Min(modifiedGroup) {
					modifiedGroup[k] = append(modifiedGroup[k], v)
					//(*s) += string(v)
					//fmt.Println(modifiedGroup[k])
					break
				}
			}
		}
	}

	for _, v := range modifiedGroup {
		r.TotalCost += float32(len(v)) * s.CostPerBook * c.Discount[len(v)-1]
	}


	r.Books = s.Books
	r.CostPerBook = s.CostPerBook

	return
}


func Max(list []int) (max int) {
	max = 0
	for i := 0; i < 5; i = i + 1 {
		counter := 0
		for _, v := range list {
			if v == i {
				counter++
			}
		}
		if counter > max {
			max = counter
		}
	}
	return
}

func Min(slice [][]int) (min int) {
	min = len(slice[0])
	for _, v := range slice {
		if len(v) < min {
			min = len(v)
		}
	}

	return
}

func Contain(list []int, i int) (res bool) {
	res = false
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return
}

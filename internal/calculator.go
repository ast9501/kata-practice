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
	//fmt.Println(s.Books)
	//fmt.Println("Max s.Books: ", Max(s.Books))
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
		//fmt.Println(float32(len(v)), " ", r.CostPerBook, " ", c.Discount[len(v)-1])
		//fmt.Println(r.TotalCost)
	}
	//fmt.Println(group)
	/*
		for i, v := range books {
			fmt.Println("i: ", i, " v: ", v)
			r.TotalCost += (float32(s.CostPerBook) * float32(len(books)-i-1) * c.Discount[len(books)-i-1]) * float32(v)
			fmt.Println(r.TotalCost)
			books = SliceDeletor(v, &books)
			fmt.Println("Books remain: ", books)
		}
	*/
	r.Books = s.Books
	r.CostPerBook = s.CostPerBook

	return
	// 2 2 4 1 0 -> 0 1 2 2 4
}

/*
func SliceDeletor(n int, i *[]int) (r []int) {
	for _, v := range *i {
		if v-n >= 0 {
			r = append(r, v-n)
		} else {
			r = append(r, 0)
		}
	}
	return
}
*/
func Max(list []int) (max int) {
	//i := 0
	max = 0
	for i := 0; i < 5; i = i + 1 {
		counter := 0
		for _, v := range list {
			if v == i {
				counter++
				//fmt.Println("v: ", v, " = i: ", i)
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

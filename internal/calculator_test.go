package internal

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	c := BasicCalculator{
		Discount: []float32{0, 0.95, 0.9, 0.8, 0.75},
	}

	s := ShoppingList{
		Books:       []int{0, 0, 1, 1, 2, 2, 3, 4},
		CostPerBook: 8.0,
		TotalCost:   0.0,
	}

	s = c.CountCost(s)
	if s.TotalCost == 51.2 {
		t.Log("BasicCalculator.CoasCount() PASS")
	} else {
		t.Error("BasicCalculator.CoasCount() FAIL, get ", s.TotalCost, ", expect: 51.2")
	}
}

func TestMax(t *testing.T) {
	l := []int{0, 0, 1, 1, 2, 2, 3, 4, 2, 5}
	if Max(l) == 3 {
		t.Log("Max() PASS")
	} else {
		t.Error("Max() FAIL, get ", Max(l), " expect 3")
	}
}

func TestMin(t *testing.T) {
	s := [][]int{[]int{0, 0, 1, 2}, []int{0, 1, 2}}
	if Min(s) == 3 {
		t.Log("Min() PASS")
	} else {
		t.Error("Min() FAIL, get ", Min(s), " expect 3")
	}
}

func TestContain(t *testing.T) {
	s := []int{0, 0, 1, 2, 3}
	if !Contain(s, 1) || Contain(s, 4) {
		if Contain(s, 1) {
			t.Error("Contain() FAIL, get ", Contain(s, 1), "expect true")
		} else {
			t.Error("Contain() FAIL, get ", Contain(s, 4), "expect false")
		}

	} else {
		t.Log("Contain() PASS")
	}
}

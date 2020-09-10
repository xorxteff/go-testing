package services

import (
	"go-testing/api/utils/sort"
)

func Sort(numbers []int) {
	sort.BubbleSort(numbers)
}

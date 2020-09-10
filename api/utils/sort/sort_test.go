package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	numbers := []int{8, 7, 4, 5, 3, 6, 8, 0}
	BubbleSort(numbers)

	fmt.Println(numbers)

	if numbers[0] != 0 {
		t.Error("first element should be 0")
	}
	if numbers[len(numbers)-1] != 8 {
		t.Error("last element should be 8")
	}
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	BubbleSort(numbers)
	fmt.Println(numbers)

	if numbers[0] != 1 {
		t.Error("first element should be 1")
	}
	if numbers[len(numbers)-1] != 4 {
		t.Error("last element should be 4")
	}

}

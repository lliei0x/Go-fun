package bubblesort

import "testing"
import "fmt"

func TestQuickSort1(t *testing.T) {
	valuse := []int{5, 4, 3, 2, 1}
	QuickSort(valuse)

	if valuse[0] != 1 || valuse[1] != 2 || valuse[2] != 3 || valuse[3] != 4 || valuse[4] != 5 {
		fmt.Println("BubbleSort() failed . Got ", valuse, "Expected 1 2 3 4 5")
	}
}

func TestQuickSort2(t *testing.T) {
	valuse := []int{5, 5, 3, 2, 1}
	QuickSort(valuse)

	if valuse[0] != 1 || valuse[1] != 2 || valuse[2] != 3 || valuse[3] != 5 || valuse[4] != 5 {
		fmt.Println("BubbleSort() failed . Got ", valuse, "Expected 1 2 3 5 5")
	}
}

func TestQuickSort3(t *testing.T) {
	valuse := []int{5}
	QuickSort(valuse)

	if valuse[0] != 5 {
		fmt.Println("BubbleSort() failed . Got ", valuse, "Expected 5")
	}
}

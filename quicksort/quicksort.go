package quicksort

// partition picks a pivot and weakly sorts the slice
// around it. It assumes the slice has a length of at
// least 2 elements.
func partition(arr []int, low int, high int) int {
	pickedPivot := high

	smallerIdx := low

	for i := low; i < high; i++ {
		if arr[i] < arr[pickedPivot] {
			aux := arr[i]
			arr[i] = arr[smallerIdx]
			arr[smallerIdx] = aux
			smallerIdx++
		}
	}

	pivotIdx := low
	for arr[pivotIdx] <= arr[pickedPivot] && pivotIdx < high {
		pivotIdx++
	}

	aux := arr[pickedPivot]
	arr[pickedPivot] = arr[pivotIdx]
	arr[pivotIdx] = aux

	return pivotIdx
}

func qs(arr []int, low int, high int) {
	if high-low < 1 {
		return
	}

	pIdx := partition(arr, low, high)

	qs(arr, low, pIdx-1)
	qs(arr, pIdx+1, high)
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

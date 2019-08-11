package sort

func MergeSort(data []int) {
	length := len(data)
	tmp := make([]int, length)
	for i := 0; i < length; i++ {
		tmp[i] = data[i]
	}
	mergeSort(data, 0, length-1, tmp)
}

func mergeSort(data []int, begin, end int, tmp []int) {
	if end <= begin {
		return
	}
	mid := (begin + end) / 2
	mergeSort(data, begin, mid, tmp)
	mergeSort(data, mid+1, end, tmp)
	merge(data, begin, mid+1, end, tmp)
}

func merge(data []int, left, right, end int, tmp []int) {
	leftEnd := right - 1
	i := left
	j := right
	index := left
	for i <= leftEnd && j <= end {
		if tmp[i] < tmp[j] {
			data[index] = tmp[i]
			index++
			i++
		} else {
			data[index] = tmp[j]
			index++
			j++
		}
	}
	for ; i <= leftEnd; i++ {
		data[index] = tmp[i]
		index++
	}
	for ; j <= end; j++ {
		data[index] = tmp[j]
		index++
	}
	for k := left; k <= end; k++ {
		tmp[k] = data[k]
	}
}

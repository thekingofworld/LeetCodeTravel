package sort

func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := (lo + hi) / 2
	mediumIndex := selectMedium(data, lo, mid, hi)
	data[lo], data[mediumIndex] = data[mediumIndex], data[lo]
	v := data[lo]
	lt := lo
	gt := hi
	i := lo + 1
	for i <= gt {
		if data[i] < v {
			data[i], data[lt] = data[lt], data[i]
			i++
			lt++
		} else if data[i] > v {
			data[i], data[gt] = data[gt], data[i]
			gt--
		} else {
			i++
		}
	}
	quickSort(data, lo, lt-1)
	quickSort(data, gt+1, hi)
}

func selectMedium(data []int, left, center, right int) int {
	if data[left] < data[center] {
		if data[center] < data[right] {
			return center
		} else if data[left] < data[right] {
			return right
		} else {
			return left
		}
	} else {
		if data[left] < data[right] {
			return left
		} else if data[center] < data[right] {
			return center
		} else {
			return right
		}
	}
}

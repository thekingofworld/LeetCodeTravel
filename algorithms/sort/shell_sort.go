package sort

func ShellSort(data []int) {
	length := len(data)
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				data[j], data[j-h] = data[j-h], data[j]
			}
		}
		h = h / 3
	}
}

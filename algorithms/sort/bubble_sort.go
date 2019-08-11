package sort

func BubbleSort(data []int) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

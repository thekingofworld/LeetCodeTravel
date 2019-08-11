package sort

func InsertSort(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		tmp := data[i]
		var j int
		for j = i; j > 0 && tmp < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		data[j] = tmp
	}
}

package sort

func HeapSort(data []int) {
	length := len(data)
	heap := make([]int, length+1)
	for i := range data {
		heap[i+1] = data[i]
	}
	for k := length / 2; k >= 1; k-- {
		sink(heap, k, length)
	}
	n := length
	for n > 0 {
		heap[1], heap[n] = heap[n], heap[1]
		n--
		sink(heap, 1, n)
	}
	for i, v := range heap[1:] {
		data[i] = v
	}
}

func sink(heap []int, k, n int) {
	for 2*k <= n {
		child := 2 * k
		if child+1 <= n && heap[child+1] > heap[child] {
			child++
		}
		if heap[k] < heap[child] {
			heap[k], heap[child] = heap[child], heap[k]
			k = child
		} else {
			break
		}
	}
}

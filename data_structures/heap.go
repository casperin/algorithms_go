package datastructures

// Implementation of a min heap where the stored value is an int.
//
// A min heap has two main functions: Push and Pop. You can push values onto it
// in any order, but pop will always return the lowest of the pushed values.
//
// This implementation has a few more convenience methods, such as Size,
// PushAll, Peek, Replace, and Meld.
type Heap struct {
	values []int
	less   func(int, int) bool
}

func NewHeap() Heap {
	return Heap{
		[]int{},
		func(i, j int) bool { return i < j },
	}
}

func NewHeapWithLess(less func(int, int) bool) Heap {
	return Heap{[]int{}, less}
}

func (h *Heap) Size() int {
	return len(h.values)
}

// Pushes a value onto the heap.
func (h *Heap) Push(value int) {
	h.values = append(h.values, value)
	h.siftUp(len(h.values) - 1)
}

// Pushes each value in the list onto the heap.
func (h *Heap) PushAll(values []int) {
	for _, value := range values {
		h.Push(value)
	}
}

// Returns the lowest value from the heap, without removing it.
func (h *Heap) Peek() (int, bool) {
	if h.Size() == 0 {
		return 0, false
	}
	return h.values[0], true
}

// Returns the lowest value on the heap and removes it.
func (h *Heap) Pop() (int, bool) {
	size := h.Size()

	if size == 0 {
		return 0, false
	}

	// The idea here, is that we take out the value we want to return (the
	// lowest value, which is in front of the list). We then move the last
	// value up in front, and let that sift down.
	min := h.values[0]                       // Save return value
	h.values[0] = h.values[size-1]           // Move value at the end up in front
	h.values = h.values[0 : len(h.values)-1] // Remove the last value
	h.siftDown(0)                            // Lift down the first value (it is unlikely to be the lowest)
	return min, true
}

// Push and Pop in one operation. The reason for doing it in one, is that we
// only need to reorganize the heap once, making the operation faster.
//
// Replace does not return the second `ok` value since we are popping after we
// just pushed a value, meaning we know for sure that the heap contains at
// least one value to pop.
func (h *Heap) Replace(value int) int {
	h.values = append(h.values, value)
	min, _ := h.Pop()
	return min
}

// Merges another heap into this heap, cosuming the other heap in the process.
func (h *Heap) Meld(k *Heap) {
	for {
		v, ok := k.Pop()
		if !ok {
			break
		}
		h.Push(v)
	}
}

// Used when pushing values onto the heap. What happens is we push them onto
// the list (at the end), and then let that value sift towards the front until
// its parent is no longer less.
func (h *Heap) siftUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := (index - 1) / 2

	if h.less(h.values[index], h.values[parentIndex]) {
		// swap this value with the parent and continue upwards from the parent
		h.values[index], h.values[parentIndex] = h.values[parentIndex], h.values[index]
		h.siftUp(parentIndex)
	}
}

// Opposite siftUp, we here have the value at the front of the list, and we
// compare it to each of its children and let it sift downwards until neither
// of its children are greater.
func (h *Heap) siftDown(index int) {
	size := h.Size()
	smallest := index
	l := index*2 + 1
	r := index*2 + 2

	if l < size && h.less(h.values[l], h.values[smallest]) {
		smallest = l // left child is smaller
	}

	if r < size && h.less(h.values[r], h.values[smallest]) {
		smallest = r // right child is smaller
	}

	// If either child is smaller, then..
	if smallest != index {
		swap(h.values, index, smallest) // swap the current index with the smallest
		h.siftDown(smallest)            // and continue down that branch
	}
}

// Helper function for siftDown
func swap(list []int, i, j int) {
	list[i], list[j] = list[j], list[i]
}

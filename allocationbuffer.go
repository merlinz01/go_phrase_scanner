package go_phrase_scanner

// A temporary buffer for reducing memory allocation overhead.
type allocationBuffer[T any] struct {
	// Function to initialize new values
	InitFunc func() T
	unitSize int
	pieces   []T
}

// Create a new pool with a size of unitSize
func newAllocationBuffer[T any](unitSize int) allocationBuffer[T] {
	pieces := make([]T, 0, unitSize)
	return allocationBuffer[T]{unitSize: unitSize, pieces: pieces}
}

// Get a pointer to a new value from the buffer.
// If the buffer is used up, allocates a new buffer.
func (p *allocationBuffer[T]) allocate() *T {
	if len(p.pieces) == p.unitSize {
		p.pieces = make([]T, 0, p.unitSize)
	}
	p.pieces = append(p.pieces, p.InitFunc())
	return &p.pieces[len(p.pieces)-1]
}

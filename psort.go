/*
Package psort implement a topological sort on a collection of partially ordered items.
*/
package psort

// Interface defines a collection of items that are partially ordered
type Interface interface {
	// Len is the number of elements in the collection
	Len() int

	// LessOrEqual reports whether the element I with index i
	// and the element J with index j are related such that i <= j
	LessOrEqual(i, j int) bool

	// Swap the elements at indexes i and j
	Swap(i, j int)
}

// Sort performs an in-place topological sort on a partially ordered collection of items.
// Returns false if the list could not be sorted, which happens if a cycle is found.
//
// The implementation is based on Tarjan's algorithm.
// See: https://en.wikipedia.org/wiki/Topological_sorting#Tarjan.27s_algorithm
func Sort(data Interface) bool {
	// A mapping of indicies to real indicies of the data
	index := make([]int, data.Len())
	for i := 0; i < len(index); i++ {
		index[i] = i
	}

	// Indicates which elements have been marked
	marked := make([]bool, len(index))

	// An index into the data collection, indicating the current location of the head of the list
	head := len(index) - 1

	var visitFn func(int) bool
	visitFn = func(n int) bool {
		if marked[n] {
			return false
		}
		marked[n] = true

		for m := 0; m < len(index); m++ {
			if m == n || index[m] > head {
				continue
			}

			// Check if there is an 'edge' from n to m
			// Equivalent to saying n <= m
			if data.LessOrEqual(index[n], index[m]) {
				if !visitFn(m) {
					return false
				}
			}
		}

		// Add n to the head
		realN := index[n]
		if realN != head {
			var headIndex int
			for headIndex = 0; index[headIndex] != head; headIndex++ {
			}
			data.Swap(realN, head)
			index[headIndex] = realN
			index[n] = head
			marked[n] = marked[headIndex]
		}
		head--
		return true
	}

	for head >= 0 {
		var i int
		for i = 0; i < len(index) && index[i] > head; i++ {
		}
		if i == len(index) {
			return false
		}
		if !visitFn(i) {
			return false
		}
	}

	return true
}

// IsSorted determines if the given collection of partialy ordered items is sorted
func IsSorted(data Interface) bool {
	length := data.Len()
	for i := 1; i < length; i++ {
		if data.LessOrEqual(i, i-1) {
			return false
		}
	}
	return true
}

// Reverse performs an in-place reversal of the order of the items in the collection
func Reverse(data Interface) {
	head := 0
	tail := data.Len() - 1
	for head < tail {
		data.Swap(head, tail)
		head++
		tail--
	}
}

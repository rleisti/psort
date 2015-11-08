Package psort implement a topological sort on a collection of partially ordered items.

The API of this package is designed to match the standard "sort" package.  You can sort any
partially ordered collection of items that implement Interface.

# Sample Usage

	const (
		A1 = iota
		A2
		A3
		B1
	)

	type item int
	type itemCollection []item

	func (data itemCollection) Len() int {
		return len(data)
	}

	func (data itemCollection) LessOrEqual(i, j int) bool {
		return lessOrEqual(data[i], data[j])
	}

	func (data itemCollection) Swap(i, j int) {
		swap := data[i]
		data[i] = data[j]
		data[j] = swap
	}

	// lessOrEqual defines a partial order on items
	func lessOrEqual(x, y item) bool {
		switch {
		case x == A1: return y == A1 || y == A2 || y == A3
		case x == A2: return y == A2 || y == A3
		case x == A3: return y == A3
		case x == B1: return y == B1 || y == B2 || y == B3
		default: return false
		}
	}

	func test() {
		var items itemCollection
		items = []item{A3, B1, A2, A1}
		Sort(items)
		// items == {A1, A2, B1, A3}
		// though any placement of B1 in the list is valid
	}

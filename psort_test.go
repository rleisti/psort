package psort

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

const (
	A1 = iota
	A2
	A3
	B1
	B2
	B3
	C1
	C2
	C3
)

type item int
type itemCollection []item

func (x item) String() string {
	switch x {
	case A1:
		return "A1"
	case A2:
		return "A2"
	case A3:
		return "A3"
	case B1:
		return "B1"
	case B2:
		return "B2"
	case B3:
		return "B3"
	case C1:
		return "C1"
	case C2:
		return "C2"
	case C3:
		return "C3"
	default:
		return "?"
	}
}

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
	case x == A1:
		return y == A1 || y == A2 || y == A3
	case x == A2:
		return y == A2 || y == A3
	case x == A3:
		return y == A3
	case x == B1:
		return y == B1 || y == B2 || y == B3
	case x == B2:
		return y == B2 || y == B3
	case x == B3:
		return y == B3
	case x == C1:
		return y == C1 || y == C2 || y == C3
	case x == C2:
		return y == C2 || y == C3
	case x == C3:
		return y == C3
	default:
		return false
	}
}

// validateSort takes a collection and returns true if it was sorted correctly
func validateSort(data itemCollection) bool {
	Sort(data)
	return IsSorted(data)
}

// validateReverse ensures that Reverse reverses the order of items in the collection
func validateReverse(data itemCollection) bool {
	original := make([]item, data.Len())
	reflect.Copy(reflect.ValueOf(original), reflect.ValueOf(data))

	Reverse(data)
	for i := 0; i < data.Len(); i++ {
		if data[i] != original[data.Len()-1-i] {
			return false
		}
	}
	return true
}

func (data itemCollection) Generate(rand *rand.Rand, size int) reflect.Value {
	collectionSize := rand.Intn(9)
	var used = make([]bool, 9)
	var collection itemCollection = make([]item, collectionSize)
	for i := 0; i < len(collection); {
		val := item(rand.Intn(9))
		if !used[val] {
			used[val] = true
			collection[i] = val
			i++
		}
	}
	return reflect.ValueOf(collection)
}

func Test_Sort(t *testing.T) {
	err := quick.Check(validateSort, nil)
	if err != nil {
		t.Error(err)
	}
}

func Test_Reverse(t *testing.T) {
	err := quick.Check(validateReverse, nil)
	if err != nil {
		t.Error(err)
	}
}

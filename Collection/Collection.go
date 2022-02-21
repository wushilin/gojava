package Collection

import (
	"fmt"
	"reflect"

	stream "github.com/wushilin/stream"
)

// Iterator for Collections
type Iterator[T any] interface {
	// Iterator must be a stream.Iterator, which defines Next() T, bool
	stream.Iterator[T]

	// Remove last returned entry (can only be called once for every Next() call)
	Remove()

	// Set the value of current entry to the new vale, returns the old value
	Set(T) T
}

// Defines a function that can test equality of two variable of the same type
// If you don't want to implement, a collection.DefaultEqualizer[T]() is provided, which
// uses reflect.DeepEquals(v1, v2 T) for simplicity
type Equalizer[T any] func(T, T) bool

// Defines common APIs a Collection should support
type Visitor[T any] func(what T) (shouldContinue bool)
type Collection[T any] interface {
	// Visit each elements with Visitor. When visitor returns false, it stops
	// Returns the number of elements visited.
	ForEach(visitor Visitor[T]) int

	// Adds a new element to the tail of the list
	// Returns whether the add was successful
	Add(element T) (added bool)

	// Adds all elements in the collection
	// Returns the number of elements added
	AddAll(elements Collection[T]) (addedCount int)

	// Test whether this collection contains the element
	// Tests equality with the specified equalizer
	// Return true if the element is found
	ContainsFunc(what T, equals Equalizer[T]) (exists bool)

	// Test whether this collection contains the element
	// Tests equality with the default equalizer
	// Return true if the element is found
	Contains(data T) (exist bool)

	// Tests whether if the collection is empty (no elements)
	IsEmpty() (isEmpty bool)

	// Returns iterator over the list
	Iterator() (iterator Iterator[T])

	// Remove all with equals tester returns number of elements removed
	RemoveAllFunc(collection Collection[T], equals Equalizer[T]) (numberOfItemsRemoved int)

	// Remove all with the default equals tester returns number of elements removed
	RemoveAll(collection Collection[T]) (numberOfItemsRemoved int)

	// Retain all elements in the current collection if it is found in the parameter collection
	// Returns number of elements removed
	RetainAllFunc(collection Collection[T], equals Equalizer[T]) (numberOfItemsRemoved int)

	// Retain all elements in the current collection if it is found in the parameter collection
	// Returns number of elements removed
	RetainAll(collection Collection[T]) (numberOfItemsRemoved int)

	// Get size of the collection
	Size() (size int)

	// Convert collection to an array
	ToArray() (array []T)

	// Convert collection to a stream
	Stream() (stream stream.Stream[T])

	// Removes all elements in the collection, returns the number of items removed
	Clear() (numberOfItemsRemoved int)
}

// Reflection.DeepEqual
func EqualsTester[T any](v1, v2 T) bool {
	return reflect.DeepEqual(v1, v2)
}

// Default equalizer that uses reflection.DeepEquals
func DefaultEqualizer[T any]() Equalizer[T] {
	return EqualsTester[T]
}

// Visit each item in iterator with visitor function.
// Stop when visitor function returns false, or iterator is fully traversed
func ForEach[T any](iter Iterator[T], visitor Visitor[T]) int {
	count := 0
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		count++
		if !visitor(next) {
			break
		}
	}
	return count
}

// Convert iterator to fixed size array
func ToArray[T any](size int, iter Iterator[T]) []T {
	result := make([]T, size)
	index := 0
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		result[index] = next
		index++
	}

	return result
}

// Prints collection to STDOUT
func PrintCollection[T any](list Collection[T]) {
	fmt.Print("[")
	count := 0
	list.ForEach(func(i T) bool {
		if count == list.Size()-1 {
			fmt.Printf("%+v", i)
		} else {
			fmt.Printf("%+v,", i)
		}
		count++
		return true
	})
	fmt.Println("]")
}

// Add elements to the given collections
func AddElementsTo[T any](what Collection[T], arg ...T) {
	for _, obj := range arg {
		what.Add(obj)
	}
}

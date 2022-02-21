package List

import (
	coll "github.com/wushilin/gojava/Collection"
)

// Defines requirement for a list
type List[T any] interface {
	// List must be a collection
	coll.Collection[T]

	//Add an element at specific location. It moves the element at the current location to the right
	// [1,2,3,4,5].AddAt(0, 12) => [12,1,2,3,4,5]
	AddAt(index int, element T) (added bool)

	// Add all elements at specific location, moves the current elements to the right
	AddAllAt(index int, elements coll.Collection[T]) (added int)

	// Get item at index
	Get(index int) (item T)

	// Get the first Index of what with equal tester
	IndexOfFunc(what T, equalizer coll.Equalizer[T]) (index int)

	// Get the first Index of what with default equal tester
	IndexOf(what T) (index int)

	// Get the last Index of what with equal tester
	LastIndexOfFunc(what T, equaliser coll.Equalizer[T]) (index int)

	// Get the last Index of what with default equal tester
	LastIndexOf(what T) (index int)

	// Set the element at index to new value newValue, returns the old value at that position
	Set(index int, newValue T) (oldValue T)

	// Remove first element that equals data with specified equalizer
	// Returns if any item was removed
	RemoveFirstFunc(data T, equals coll.Equalizer[T]) (anyItemRemoved bool)

	// Same as RemoveFirstFunc, but with default equals tester
	RemoveFirst(v T) (anyItemRemoved bool)

	// Remove at index
	RemoveAt(index int) T

	// Make a copy of list as sublist, from fromIndexIncluded, to endIndexExcluded
	CopySubList(fromIndexIncluded int, endIndexExcluded int) (newList List[T])

	// Make a copy of the list
	Copy() (newList List[T])

	// Reverse a list and return as new list
	Reverse() (newList List[T])
}

func LinkedListOf[T any](arg ...T) *LinkedList[T] {
	list := NewLinkedList[T]()
	coll.AddElementsTo[T](list, arg...)
	return list
}

func ArrayListOf[T any](arg ...T) *ArrayList[T] {
	list := NewArrayList[T]()
	coll.AddElementsTo[T](list, arg...)
	return list
}

func ListEquals[T any](list1, list2 List[T], equalFunc coll.Equalizer[T]) bool {
	if list1.Size() != list2.Size() {
		return false
	}

	iter1 := list1.Iterator()
	iter2 := list2.Iterator()
	for next, ok := iter1.Next(); ok; next, ok = iter1.Next() {
		next2, ok2 := iter2.Next()
		if !ok2 {
			return false
		}
		if !equalFunc(next, next2) {
			return false
		}
	}
	return true
}

func FindItem[T any](iter coll.Iterator[T], data T, equals coll.Equalizer[T]) int {
	return FindItemFlag(iter, data, equals, true)
}

func FindItemFlag[T any](iter coll.Iterator[T], data T, equals coll.Equalizer[T], breakOnFirst bool) int {
	index := -1
	theIndex := 0
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if equals(next, data) {
			index = theIndex
			if breakOnFirst {
				break
			}
		}
		theIndex++
	}
	return index
}

func FindLastItem[T any](iter coll.Iterator[T], data T, equals coll.Equalizer[T]) int {
	return FindItemFlag(iter, data, equals, false)
}

func RemoveAll[T any](src coll.Collection[T], collection coll.Collection[T]) int {
	return RemoveAllFunc(src, collection, coll.DefaultEqualizer[T]())
}

func RemoveAllFunc[T any](src coll.Collection[T], what coll.Collection[T], equals coll.Equalizer[T]) int {
	count := 0
	what.ForEach(
		func(toDelete T) bool {
			iter := src.Iterator()
			for next, ok := iter.Next(); ok; next, ok = iter.Next() {
				if equals(next, toDelete) {
					iter.Remove()
					count++
				}
			}
			return true
		})
	return count
}

func RetainAll[T any](src coll.Collection[T], what coll.Collection[T]) int {
	return RetainAllFunc[T](src, what, coll.DefaultEqualizer[T]())
}

func RetainAllFunc[T any](src coll.Collection[T], what coll.Collection[T], equals coll.Equalizer[T]) int {
	removedCount := 0
	iter := src.Iterator()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		iter2 := what.Iterator()
		found := false
		for next2, ok2 := iter2.Next(); ok2; next2, ok2 = iter2.Next() {
			if equals(next, next2) {
				found = true
				break
			}
		}
		if !found {
			iter.Remove()
			removedCount++
		}
	}
	return removedCount
}

func RemoveFirstFunc[T any](iter coll.Iterator[T], data T, equals coll.Equalizer[T]) bool {
	return true
}

func RemoveAt[T any](iter coll.Iterator[T], index int) T {
	count := 0
	var result T
	found := false
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if count == index {
			found = true
			iter.Remove()
			result = item
			break
		}
		count++
	}

	if found {
		return result
	}
	panic("Invalid index")
}

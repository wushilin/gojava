package List

import (
	"fmt"
	"testing"

	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/gojava/common"
	"github.com/wushilin/stream"
)

func TestArrayList2(t *testing.T) {
	list := ArrayListOf(1, 2, 3, 4, 5)
	coll.PrintCollection[int](list)
	list.RemoveAt(4)
	coll.PrintCollection[int](list)
	iter := list.Iterator()
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		fmt.Println("Found", item)
		iter.Remove()
		fmt.Println("Remaining", list.Size())
	}
}
func T2estArrayList(t *testing.T) {
	size := 10000
	newlist := NewArrayList[int]()
	for i := 0; i < size; i++ {
		newlist.Add(i)
		common.AssertEq(t, newlist.Size(), i+1)
	}

	sum := 0
	newlist.ForEach(
		func(i int) bool {
			sum += i
			return true
		})

	common.AssertEq(t, size, newlist.Size())
	common.AssertEq(t, sum, (size*(size-1))/2)

	list1 := ArrayListOf(1, 2, 3, 4, 5)
	list2 := ArrayListOf(6, 7, 8, 9, 10)
	list1.AddAll(list2)

	list3 := ArrayListOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	coll.PrintCollection[int](list1)
	common.AssertTrue(t, ListEquals[int](list1, list3, func(i, j int) bool {
		return i == j
	}))

	list1.Clear()
	common.AssertEq(t, list1.Size(), 0)
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)
	list1.Add(5)
	list1.AddAllAt(2, list2)
	coll.PrintCollection[int](list1)
	list1.AddAt(5, 555)
	common.AssertEq(t, list1.Get(5), 555)
	list1.AddAt(0, 2077)
	common.AssertEq(t, list1.Get(0), 2077)
	list1.Add(7882)
	common.AssertEq(t, list1.Get(list1.Size()-1), 7882)
	common.AssertTrue(t, list1.ContainsFunc(7882, func(i, j int) bool {
		return i == j
	}))
	fmt.Println("Testing if it contains 7882 using default equals")
	common.AssertTrue(t, list1.Contains(7882))

	fmt.Println("Testing if it contains 7883 using default equals")
	common.AssertFalse(t, list1.Contains(7883))
	coll.PrintCollection[int](list1)
	iter := newlist.Iterator()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if next%10 != 0 {
			iter.Remove()
		} else {
			iter.Set(next + 7)
		}
	}
	fmt.Println("Size is ", newlist.Size())

	list1 = ArrayListOf(1, 2, 3, 4, 5)
	list2 = ArrayListOf(3, 4, 5, 6, 7)
	list1.RetainAll(list2)

	common.AssertTrue(t, ListEquals[int](list1, ArrayListOf(3, 4, 5), coll.DefaultEqualizer[int]()))
	list1 = ArrayListOf(1, 2, 3, 4, 5)
	list1.RemoveAll(list2)
	common.AssertTrue(t, ListEquals[int](list1, ArrayListOf(1, 2), coll.DefaultEqualizer[int]()))

	newlist.ForEach(func(i int) bool {
		fmt.Println(i)
		return true
	})

	lists := ArrayListOf[string]()
	lists.Add("world")
	lists.AddAt(0, "hello")
	coll.PrintCollection[string](lists)
	common.AssertEq(t, lists.Size(), 2)

	lists.AddAt(1, "you")
	lists2 := lists.CopySubList(0, lists.Size())

	lists.AddAllAt(3, lists2)

	listsbak := lists.CopySubList(0, lists.Size())

	lists = listsbak.CopySubList(0, listsbak.Size()).(*ArrayList[string])

	common.AssertEq(t, "you", lists.Get(4))
	fmt.Println("Last index of you", lists.LastIndexOf("you"), lists.Size())
	fmt.Println("First index of you", lists.IndexOf("you"), lists.Size())
	coll.PrintCollection[string](lists)
	stream.
		Map(lists.Stream(),
			func(i string) int {
				return len(i)
			}).
		Filter(func(i int) bool {
			return i >= 4
		}).
		Each(
			func(i int) {
				fmt.Println(i)
			})

	lists = ArrayListOf("1", "2", "3")
	common.AssertTrue(t, ListEquals[string](lists.Reverse(), ArrayListOf("3", "2", "1"), coll.DefaultEqualizer[string]()))
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func T1estList1(t *testing.T) {
	size := 10000
	newlist := NewLinkedList[int]()
	for i := 0; i < size; i++ {
		newlist.Add(i)
		common.AssertEq(t, newlist.Size(), i+1)
	}

	sum := 0
	newlist.ForEach(
		func(i int) bool {
			sum += i
			return true
		})

	common.AssertEq(t, size, newlist.Size())
	common.AssertEq(t, sum, (size*(size-1))/2)

	list1 := LinkedListOf(1, 2, 3, 4, 5)
	list2 := LinkedListOf(6, 7, 8, 9, 10)
	list1.AddAll(list2)

	list3 := LinkedListOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	coll.PrintCollection[int](list1)
	common.AssertTrue(t, ListEquals[int](list1, list3, func(i, j int) bool {
		return i == j
	}))

	list1.Clear()
	common.AssertEq(t, list1.Size(), 0)
	list1.Add(1)
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)
	list1.Add(5)
	list1.AddAllAt(2, list2)
	coll.PrintCollection[int](list1)
	list1.AddAt(5, 555)
	common.AssertEq(t, list1.Get(5), 555)
	list1.AddHead(2077)
	common.AssertEq(t, list1.Get(0), 2077)
	list1.AddTail(7882)
	common.AssertEq(t, list1.Get(list1.Size()-1), 7882)
	common.AssertTrue(t, list1.ContainsFunc(7882, func(i, j int) bool {
		return i == j
	}))
	fmt.Println("Testing if it contains 7882 using default equals")
	common.AssertTrue(t, list1.Contains(7882))

	fmt.Println("Testing if it contains 7883 using default equals")
	common.AssertFalse(t, list1.Contains(7883))
	coll.PrintCollection[int](list1)
	iter := newlist.Iterator()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if next%10 != 0 {
			iter.Remove()
		} else {
			iter.Set(next + 7)
		}
	}
	fmt.Println("Size is ", newlist.Size())

	list1 = LinkedListOf(1, 2, 3, 4, 5)
	list2 = LinkedListOf(3, 4, 5, 6, 7)
	list1.RetainAll(list2)

	common.AssertTrue(t, ListEquals[int](list1, LinkedListOf(3, 4, 5), coll.DefaultEqualizer[int]()))
	list1 = LinkedListOf(1, 2, 3, 4, 5)
	list1.RemoveAll(list2)
	common.AssertTrue(t, ListEquals[int](list1, LinkedListOf(1, 2), coll.DefaultEqualizer[int]()))

	newlist.ForEach(func(i int) bool {
		fmt.Println(i)
		return true
	})

	lists := LinkedListOf[string]()
	lists.AddTail("world")
	lists.AddHead("hello")
	coll.PrintCollection[string](lists)
	common.AssertEq(t, lists.Size(), 2)

	lists.AddAt(1, "you")
	lists2 := lists.CopySubList(0, lists.Size())

	lists.AddAllAt(3, lists2)

	listsbak := lists.CopySubList(0, lists.Size())
	for ok, elem := lists.RemoveHead(); ok; ok, elem = lists.RemoveHead() {
		fmt.Println("Removed head", elem, "remaining", lists.Size())
	}

	lists = listsbak.CopySubList(0, listsbak.Size()).(*LinkedList[string])

	for ok, elem := lists.RemoveTail(); ok; ok, elem = lists.RemoveTail() {
		fmt.Println("Removed tail", elem, "remaining", lists.Size())
	}

	lists = listsbak.CopySubList(0, listsbak.Size()).(*LinkedList[string])

	common.AssertEq(t, "you", lists.Get(4))
	fmt.Println("Last index of you", lists.LastIndexOf("you"), lists.Size())
	fmt.Println("First index of you", lists.IndexOf("you"), lists.Size())
	coll.PrintCollection[string](lists)
	stream.
		Map(lists.Stream(),
			func(i string) int {
				return len(i)
			}).
		Filter(func(i int) bool {
			return i >= 4
		}).
		Each(
			func(i int) {
				fmt.Println(i)
			})

	lists = LinkedListOf("1", "2", "3")
	common.AssertTrue(t, ListEquals[string](lists.Reverse(), LinkedListOf("3", "2", "1"), coll.DefaultEqualizer[string]()))
}

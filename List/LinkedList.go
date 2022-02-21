package List

import (
	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/stream"
)

type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}

type LinkedList[T any] struct {
	head       *linkedListNode[T]
	tail       *linkedListNode[T]
	size       int
	generation int
}

func (v *LinkedList[T]) Size() int {
	return v.size
}

func (v *LinkedList[T]) removeNode(node *linkedListNode[T]) {
	if node == nil {
		panic("Can't remove nil node")
	}
	if v.size <= 1 {
		v.head = nil
		v.tail = nil
		v.size = 0
		return
	}
	if node == v.head {
		newhead := node.next
		node.next = nil
		v.head = newhead
		newhead.prev = nil
		v.size--
		return
	}
	if node == v.tail {
		newtail := node.prev
		node.prev = nil
		v.tail = newtail
		newtail.next = nil
		v.size--
		return
	}

	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.prev = nil
	node.next = nil
	v.size--
}
func (v *LinkedList[T]) Add(data T) bool {
	return v.AddTail(data)
}

func (v *LinkedList[T]) AddTail(data T) bool {
	defer v.applyMod()
	if v.size == 0 {
		v.head = &linkedListNode[T]{data: data, next: nil, prev: nil}
		v.tail = v.head
		v.size = 1
		return true
	}
	v.addAfter(v.tail, data)
	return true
}

func (v *LinkedList[T]) applyMod() {
	v.generation++
}

func (v *LinkedList[T]) RemoveHead() (bool, T) {
	if v.size == 0 {
		var zv T
		return false, zv
	}

	defer v.applyMod()
	data := v.head.data
	v.removeNode(v.head)
	return true, data
}

func (v *LinkedList[T]) RemoveTail() (bool, T) {
	if v.size == 0 {
		var zv T
		return false, zv
	}

	defer v.applyMod()
	data := v.tail.data
	v.removeNode(v.tail)
	return true, data
}

func (v *LinkedList[T]) AddHead(data T) bool {
	if v.size == 0 {
		v.AddTail(data)
		return true
	}
	defer v.applyMod()
	v.addBefore(v.head, data)
	return true
}

func (v *LinkedList[T]) ForEach(visitor coll.Visitor[T]) int {
	return coll.ForEach(v.Iterator(), visitor)
}

func (v *LinkedList[T]) addBefore(node *linkedListNode[T], data T) {
	newNode := &linkedListNode[T]{data: data, prev: nil, next: node}
	before := node.prev
	node.prev = newNode
	if before == nil {
		v.head = newNode
	} else {
		before.next = newNode
		newNode.prev = before
	}
	v.size++
}
func (v *LinkedList[T]) addAfter(node *linkedListNode[T], data T) {
	newNode := &linkedListNode[T]{data: data, prev: node, next: nil}
	after := node.next
	node.next = newNode
	if after == nil {
		v.tail = newNode
	} else {
		after.prev = newNode
		newNode.next = after
	}
	v.size++
}

func (v *LinkedList[T]) AddAll(elements coll.Collection[T]) int {
	iter := elements.Iterator()
	count := 0
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if v.Add(next) {
			count++
		}
	}
	return count
}

func (v *LinkedList[T]) nodeAt(at int) *linkedListNode[T] {
	if at < 0 || at > v.size {
		panic("Index Out Of Bound")
	}

	index := 0
	node := v.head
	for ; index < at; index++ {
		node = node.next
	}
	return node
}

func (v *LinkedList[T]) AddAllAt(at int, elements coll.Collection[T]) int {
	iter := elements.Iterator()
	count := 0

	if at == v.size {
		return v.AddAll(elements)
	} else {
		node := v.nodeAt(at)
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			v.addBefore(node, next)
			count++
		}
	}
	if count > 0 {
		defer v.applyMod()
	}
	return count
}

func (v *LinkedList[T]) AddAt(index int, data T) bool {
	if index == v.size {
		return v.Add(data)
	}
	defer v.applyMod()
	node := v.nodeAt(index)
	v.addBefore(node, data)

	return true
}

func (v *LinkedList[T]) Clear() int {
	defer v.applyMod()
	oldSize := v.size
	v.head = nil
	v.tail = nil
	v.size = 0
	return oldSize
}

func (v *LinkedList[T]) Contains(data T) bool {
	return v.ContainsFunc(data, coll.DefaultEqualizer[T]())
}

func (v *LinkedList[T]) ContainsFunc(data T, equals coll.Equalizer[T]) bool {
	iterator := v.Iterator()
	index := FindItem(iterator, data, equals)
	return index != -1
}

func (v *LinkedList[T]) Get(index int) T {
	node := v.nodeAt(index)
	return node.data
}

func (v *LinkedList[T]) IndexOf(data T) int {
	return v.IndexOfFunc(data, coll.DefaultEqualizer[T]())
}

func (v *LinkedList[T]) IndexOfFunc(data T, equals coll.Equalizer[T]) int {
	index := -1
	count := 0
	v.ForEach(func(arg T) bool {
		if equals(data, arg) {
			index = count
			return false
		} else {
			count++
			return true
		}
	})
	return index
}

func (v *LinkedList[T]) beginRangeCheck(index int) {
	if index < 0 || index > v.size-1 {
		panic("Begin Index Out Of bound")
	}
}

func (v *LinkedList[T]) endRangeCheck(index int) {
	if index < 0 || index > v.size {
		panic("End Index Out Of Bound")
	}
}
func (v *LinkedList[T]) CopySubList(start, end int) List[T] {
	v.beginRangeCheck(start)
	v.endRangeCheck(end)

	length := end - start
	if length < 0 {
		panic("Invalid start & end combination")
	}
	result := NewLinkedList[T]()

	node := v.nodeAt(start)
	for result.Size() < length {
		result.Add(node.data)
		node = node.next
	}
	return result
}

func (v *LinkedList[T]) Copy() List[T] {
	return v.CopySubList(0, v.Size())
}

func (v *LinkedList[T]) Reverse() List[T] {
	result := NewLinkedList[T]()
	v.ForEach(func(i T) bool {
		result.AddHead(i)
		return true
	})
	return result
}

func (v *LinkedList[T]) IsEmpty() bool {
	return v.size == 0
}

type LinkedListIterator[T any] struct {
	Src        *LinkedList[T]
	Current    *linkedListNode[T]
	last       *linkedListNode[T]
	generation int
}

func (v *LinkedList[T]) Iterator() coll.Iterator[T] {
	return &LinkedListIterator[T]{v, v.head, nil, v.generation}
}

func (v *LinkedListIterator[T]) checkMod() {
	if v.generation != v.Src.generation {
		panic("Concurrent modification")
	}
}
func (v *LinkedListIterator[T]) Next() (T, bool) {
	v.checkMod()
	if v.Current == nil {
		var zv T
		return zv, false
	}
	result := v.Current.data
	v.last = v.Current
	v.Current = v.Current.next
	return result, true
}

func (v *LinkedListIterator[T]) Remove() {
	v.checkMod()
	defer v.applyMod()
	if v.last != nil {
		v.Src.removeNode(v.last)
		v.last = nil
	} else {
		panic("Don't call Remove when you have not read. And don't call remove twice")
	}
}

func (v *LinkedListIterator[T]) applyMod() {
	v.generation++
	v.Src.generation = v.generation
}

func (v *LinkedListIterator[T]) Set(data T) T {
	v.checkMod()
	defer v.applyMod()
	result := v.last.data
	v.last.data = data
	v.generation++
	v.Src.generation = v.generation
	return result
}

func (v *LinkedList[T]) RemoveAllFunc(what coll.Collection[T], equals coll.Equalizer[T]) int {
	count := 0
	what.ForEach(
		func(toDelete T) bool {
			iter := v.Iterator()
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

func (v *LinkedList[T]) RetainAllFunc(what coll.Collection[T], equals coll.Equalizer[T]) int {
	removedCount := 0
	iter := v.Iterator()
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

func (v *LinkedList[T]) Set(index int, data T) (oldData T) {
	defer v.applyMod()
	node := v.nodeAt(index)
	oldData = node.data
	node.data = data
	return oldData
}
func (v *LinkedList[T]) RemoveFirst(data T) bool {
	return v.RemoveFirstFunc(data, coll.DefaultEqualizer[T]())
}

func (v *LinkedList[T]) RemoveAll(collection coll.Collection[T]) int {
	return v.RemoveAllFunc(collection, coll.DefaultEqualizer[T]())
}
func (v *LinkedList[T]) RetainAll(collection coll.Collection[T]) int {
	return v.RetainAllFunc(collection, coll.DefaultEqualizer[T]())
}

func (v *LinkedList[T]) Stream() stream.Stream[T] {
	return stream.FromIterator[T](v.Iterator())
}

func (v *LinkedList[T]) RemoveFirstFunc(data T, equals coll.Equalizer[T]) bool {
	return RemoveFirstFunc(v.Iterator(), data, equals)
}

func (v *LinkedList[T]) ToArray() []T {
	return coll.ToArray(v.Size(), v.Iterator())
}

func (v *LinkedList[T]) RemoveAt(index int) T {
	return RemoveAt(v.Iterator(), index)
}

func (v *LinkedList[T]) LastIndexOf(what T) int {
	return v.LastIndexOfFunc(what, coll.DefaultEqualizer[T]())
}

func (v *LinkedList[T]) LastIndexOfFunc(what T, equals coll.Equalizer[T]) int {
	pointer := v.size - 1
	for node := v.tail; node != nil; node = node.prev {
		if equals(what, node.data) {
			return pointer
		} else {
			pointer--
		}
	}
	return -1
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

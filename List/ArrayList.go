package List

import (
	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/stream"
)

type ArrayList[T any] struct {
	buffer     []T
	generation int
	length     int
}

func (v *ArrayList[T]) Size() int {
	return v.length
}

func (v *ArrayList[T]) expand() {
	v.ensureCapacity(len(v.buffer) + len(v.buffer)/4 + 5)
}

func (v *ArrayList[T]) ensureCapacity(target int) {
	if len(v.buffer) >= target {
		return
	}
	newBuffer := make([]T, target)
	for i := 0; i < v.length; i++ {
		newBuffer[i] = v.buffer[i]
	}

	v.buffer = newBuffer
}

func (v *ArrayList[T]) indexCheck(what int) {
	if what < 0 || what >= v.length {
		panic("Index Out of Bound")
	}
}

func (v *ArrayList[T]) applyMod() {
	v.generation++
}

func (v *ArrayList[T]) Get(index int) T {
	v.indexCheck(index)
	return v.buffer[index]
}

func (v *ArrayList[T]) Add(what T) bool {
	defer v.applyMod()
	v.ensureCapacity(v.Size() + 1)
	v.buffer[v.length] = what
	v.length++
	return true
}

func (v *ArrayList[T]) AddAt(index int, what T) bool {
	defer v.applyMod()
	v.ensureCapacity(v.length + 1)
	if index == v.length {
		v.buffer[v.length] = what
		v.length++
		return true
	}

	v.shiftRight(index, 1)
	v.buffer[index] = what
	v.length++
	return true

}
func (v *ArrayList[T]) shiftLeft(fromWhere int, slots int) {
	// 1,2,3,4,5 shiftLeft 2,2
	if slots > fromWhere {
		panic("Invalid shift left arguments")
	}
	for i := fromWhere; i < v.length; i++ {
		v.buffer[i-slots] = v.buffer[i]
	}
}
func (v *ArrayList[T]) shiftRight(fromWhere int, slots int) {
	// 1,2,3,4,5, shiftright 2, 3 //1,2,_,_,_,3,4,5
	tmpBuffer := make([]T, v.length-fromWhere)
	for i := 0; i < len(tmpBuffer); i++ {
		tmpBuffer[i] = v.buffer[i+fromWhere]
	}
	for index, next := range tmpBuffer {
		v.buffer[fromWhere+slots+index] = next
	}
}
func (v *ArrayList[T]) AddAll(what coll.Collection[T]) int {
	return v.AddAllAt(v.length, what)
}

func (v *ArrayList[T]) AddAllAt(index int, what coll.Collection[T]) int {
	defer v.applyMod()
	oldLength := v.length
	newLength := oldLength + what.Size()
	v.ensureCapacity(newLength)
	if index == v.length {
		// Add tail
		iter := what.Iterator()
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			v.buffer[v.length] = next
			v.length++
		}
		return what.Size()
	}

	v.shiftRight(index, what.Size())

	what.ForEach(func(i T) bool {
		v.buffer[index] = i
		index++
		return true
	})

	v.length = newLength
	return what.Size()
}

func (v *ArrayList[T]) Clear() int {
	defer v.applyMod()
	size := v.Size()
	v.length = 0
	return size
}

func (v *ArrayList[T]) Contains(data T) bool {
	return v.ContainsFunc(data, coll.DefaultEqualizer[T]())
}

func (v *ArrayList[T]) ContainsFunc(what T, equals coll.Equalizer[T]) bool {
	return FindItem(v.Iterator(), what, equals) != -1
}

func (v *ArrayList[T]) Copy() (newList List[T]) {
	return v.CopySubList(0, v.Size())
}

func (v *ArrayList[T]) CopySubList(startInclude int, endExclude int) (newList List[T]) {
	result := NewArrayList[T]()
	result.ensureCapacity(endExclude - startInclude)

	for i := startInclude; i < endExclude; i++ {
		result.Add(v.Get(i))
	}
	return result
}

func (v *ArrayList[T]) ForEach(visitor coll.Visitor[T]) int {
	return coll.ForEach(v.Iterator(), visitor)
}

func (v *ArrayList[T]) Set(index int, data T) T {
	defer v.applyMod()
	v.indexCheck(index)
	old := v.Get(index)
	v.buffer[index] = data
	return old
}

func (v *ArrayList[T]) Iterator() coll.Iterator[T] {
	return &ArrayListIterator[T]{Src: v, currentIndex: 0, lastReturnedIndex: -1, generation: v.generation}
}

func (v *ArrayList[T]) IndexOf(data T) int {
	return v.IndexOfFunc(data, coll.DefaultEqualizer[T]())
}

func (v *ArrayList[T]) IndexOfFunc(data T, equals coll.Equalizer[T]) int {
	return FindItem(v.Iterator(), data, equals)
}

func (v *ArrayList[T]) IsEmpty() bool {
	return v.length == 0
}

func (v *ArrayList[T]) LastIndexOf(data T) int {
	return v.LastIndexOfFunc(data, coll.DefaultEqualizer[T]())
}

func (v *ArrayList[T]) LastIndexOfFunc(data T, equals coll.Equalizer[T]) int {
	return FindLastItem(v.Iterator(), data, equals)
}

func (v *ArrayList[T]) RemoveAll(collection coll.Collection[T]) int {
	return RemoveAllFunc[T](v, collection, coll.DefaultEqualizer[T]())
}

func (v *ArrayList[T]) RemoveAllFunc(what coll.Collection[T], equals coll.Equalizer[T]) int {
	return RemoveAllFunc[T](v, what, equals)
}

func (v *ArrayList[T]) RetainAll(collection coll.Collection[T]) int {
	return v.RetainAllFunc(collection, coll.DefaultEqualizer[T]())
}

func (v *ArrayList[T]) RetainAllFunc(collection coll.Collection[T], equals coll.Equalizer[T]) int {
	return RetainAllFunc[T](v, collection, equals)
}

func (v *ArrayList[T]) RemoveFirst(data T) bool {
	return v.RemoveFirstFunc(data, coll.DefaultEqualizer[T]())
}

func copyArray[T any](what []T) []T {
	result := make([]T, len(what))
	for idx, val := range what {
		result[idx] = val
	}
	return result
}

func reverseCopyArray[T any](what []T) []T {
	length := len(what)
	result := make([]T, len(what))
	for idx, val := range what {
		result[length-idx-1] = val
	}
	return result
}
func (v *ArrayList[T]) Reverse() List[T] {
	newdata := reverseCopyArray(v.buffer[:v.length])
	return &ArrayList[T]{buffer: newdata, length: v.length, generation: 0}

}
func (v *ArrayList[T]) RemoveFirstFunc(data T, equals coll.Equalizer[T]) bool {
	return RemoveFirstFunc(v.Iterator(), data, equals)
}
func (v *ArrayList[T]) Stream() stream.Stream[T] {
	return stream.FromIterator[T](v.Iterator())
}

func (v *ArrayList[T]) ToArray() []T {
	return coll.ToArray(v.Size(), v.Iterator())
}

func (v *ArrayList[T]) RemoveAt(index int) T {
	return RemoveAt(v.Iterator(), index)
}

type ArrayListIterator[T any] struct {
	Src               *ArrayList[T]
	currentIndex      int
	lastReturnedIndex int
	generation        int
}

func (v *ArrayListIterator[T]) checkMod() {
	if v.generation != v.Src.generation {
		panic("Concurrent Modification detected")
	}
}
func (v *ArrayListIterator[T]) applyMod() {
	v.generation++
	v.Src.generation = v.generation
}

func (v *ArrayListIterator[T]) Next() (val T, ok bool) {
	v.checkMod()
	if v.currentIndex < 0 || v.currentIndex >= v.Src.Size() {
		var zv T
		return zv, false
	}
	result := v.Src.Get(v.currentIndex)
	v.lastReturnedIndex = v.currentIndex
	v.currentIndex++
	return result, true
}

func (v *ArrayListIterator[T]) Remove() {
	v.checkMod()
	if v.lastReturnedIndex == -1 || v.lastReturnedIndex == v.currentIndex {
		panic("Don't call Remove when you have not read, or you have removed")
	}
	defer v.applyMod()
	v.Src.shiftLeft(v.lastReturnedIndex+1, 1)
	v.Src.length--
	v.lastReturnedIndex = -1
	v.currentIndex--
}

func (v *ArrayListIterator[T]) Set(data T) T {
	v.checkMod()
	if v.lastReturnedIndex == -1 {
		panic("Don't call Set when you have not read, or you have removed")
	}
	defer v.applyMod()
	return v.Src.Set(v.lastReturnedIndex, data)
}

// Return new empty ArrayList[T]
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{buffer: make([]T, 10), length: 0}
}

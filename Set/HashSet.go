package Set

import (
	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/stream"
)

type Set[T comparable] interface {
	coll.Collection[T]
}

type HashSet[T comparable] struct {
	data       map[T]any
	generation int
}

func (v *HashSet[T]) Contains(what T) bool {
	_, ok := v.data[what]
	return ok
}

func (v *HashSet[T]) ContainsFunc(what T, equals coll.Equalizer[T]) bool {
	return v.Contains(what)
}

func (v *HashSet[T]) ForEach(visitor coll.Visitor[T]) int {
	return coll.ForEach(v.Iterator(), visitor)
}

func (v *HashSet[T]) applyMod() {
	v.generation++
}

func (v *HashSet[T]) Add(data T) bool {
	defer v.applyMod()
	if v.Contains(data) {
		return false
	}
	v.data[data] = true
	return true
}

func (v *HashSet[T]) Clear() int {
	defer v.applyMod()
	old := v.Size()
	v.data = make(map[T]any)
	return old
}

func (v *HashSet[T]) IsEmpty() bool {
	return v.Size() == 0
}

func (v *HashSet[T]) AddAll(data coll.Collection[T]) int {
	count := 0
	data.ForEach(func(i T) bool {
		if v.Add(i) {
			count++
		}
		return true
	})
	return count
}

func (v *HashSet[T]) Size() int {
	return len(v.data)
}

func (v *HashSet[T]) Iterator() coll.Iterator[T] {
	return NewIteratorFor(v)
}

func (v *HashSet[T]) ContainsAll(what coll.Collection[T]) bool {
	containsAll := true
	what.ForEach(func(i T) bool {
		if !v.Contains(i) {
			containsAll = false
			return false
		}
		return true
	})
	return containsAll
}

func (v *HashSet[T]) Remove(what T) bool {
	if v.Contains(what) {
		defer v.applyMod()
		delete(v.data, what)
		return true
	}
	return false
}
func (v *HashSet[T]) RemoveAll(what coll.Collection[T]) int {
	count := 0
	what.ForEach(func(key T) bool {
		if v.Remove(key) {
			count++
		}
		return true
	})
	return count
}

func (v *HashSet[T]) RemoveAllFunc(what coll.Collection[T], equals coll.Equalizer[T]) int {
	return v.RemoveAll(what)
}

func (v *HashSet[T]) RetainAllFunc(what coll.Collection[T], equals coll.Equalizer[T]) int {
	return v.RetainAll(what)
}

func (v *HashSet[T]) RetainAll(what coll.Collection[T]) int {
	iter := v.Iterator()
	count := 0
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		if !what.Contains(item) {
			iter.Remove()
			count++
		}
	}
	return count
}

func (v *HashSet[T]) Stream() stream.Stream[T] {
	return stream.FromIterator[T](v.Iterator())
}

func (v *HashSet[T]) ToArray() []T {
	return coll.ToArray(v.Size(), v.Iterator())
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: make(map[T]any), generation: 0}
}

func HashSetOf[T comparable](args ...T) *HashSet[T] {
	result := NewHashSet[T]()
	coll.AddElementsTo[T](result, args...)
	return result
}

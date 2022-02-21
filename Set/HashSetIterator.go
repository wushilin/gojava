package Set

import coll "github.com/wushilin/gojava/Collection"

type HashSetIterator[T comparable] struct {
	Src          *HashSet[T]
	generation   int
	currentIndex int
	lastIndex    int
	keys         []T
}

func (v *HashSetIterator[T]) applyMod() {
	v.generation++
	v.Src.generation = v.generation
}

func (v *HashSetIterator[T]) checkMod() {
	if v.generation != v.Src.generation {
		panic("Concurrent modification")
	}
}

func (v *HashSetIterator[T]) indexCheck(index int) {
	if index < 0 || index >= len(v.keys) {
		panic("Index out of bound")
	}
}

func (v *HashSetIterator[T]) Next() (result T, ok bool) {
	v.checkMod()
	if v.currentIndex >= len(v.keys) {
		var zv T
		return zv, false
	}
	v.indexCheck(v.currentIndex)
	result = v.keys[v.currentIndex]
	v.currentIndex++
	v.lastIndex = v.currentIndex - 1
	return result, true
}

func (v *HashSetIterator[T]) Remove() {
	v.checkMod()
	defer v.applyMod()
	if v.lastIndex == -1 || v.lastIndex == v.currentIndex {
		panic("Don't call remove before reading, and don't remove twice")
	}
	lastKey := v.keys[v.lastIndex]
	v.lastIndex = -1
	delete(v.Src.data, lastKey)
}

func (v *HashSetIterator[T]) Set(data T) T {
	v.checkMod()
	defer v.applyMod()
	lastIndex := v.lastIndex
	if lastIndex == -1 {
		lastIndex = v.currentIndex - 1
	}
	if lastIndex == -1 || lastIndex == v.currentIndex {
		panic("Don't call set before reading")
	}
	lastKey := v.keys[lastIndex]
	delete(v.Src.data, lastKey)
	v.Src.data[data] = true
	v.keys[lastIndex] = data
	return lastKey
}

func readKeys[T comparable](mp map[T]any) []T {
	result := make([]T, len(mp))
	index := 0
	for key := range mp {
		result[index] = key
		index++
	}
	return result
}

func NewIteratorFor[T comparable](v *HashSet[T]) coll.Iterator[T] {
	return &HashSetIterator[T]{Src: v, generation: v.generation, currentIndex: 0, lastIndex: -1, keys: readKeys(v.data)}
}

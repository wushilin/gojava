package Map

import coll "github.com/wushilin/gojava/Collection"

type HashMapIterator[K comparable, V any] struct {
	Src          *HashMap[K, V]
	generation   int
	currentIndex int
	lastIndex    int
	keys         []K
}

func (v *HashMapIterator[K, V]) applyMod() {
	v.generation++
	v.Src.generation = v.generation
}

func (v *HashMapIterator[K, V]) checkMod() {
	if v.generation != v.Src.generation {
		panic("Concurrent modification")
	}
}

func (v *HashMapIterator[K, V]) indexCheck(index int) {
	if index < 0 || index >= len(v.keys) {
		panic("Index out of bound")
	}
}

func (v *HashMapIterator[K, V]) Next() (result KV[K, V], ok bool) {
	v.checkMod()
	if v.currentIndex >= len(v.keys) {
		var zv KV[K, V]
		return zv, false
	}
	v.indexCheck(v.currentIndex)
	resultKey := v.keys[v.currentIndex]
	resultValue, ok := v.Src.Get(resultKey)

	v.currentIndex++
	v.lastIndex = v.currentIndex - 1
	return KVOf(resultKey, resultValue), true
}

func (v *HashMapIterator[K, V]) Remove() {
	v.checkMod()
	defer v.applyMod()
	if v.lastIndex == -1 || v.lastIndex == v.currentIndex {
		panic("Don't call remove before reading, and don't remove twice")
	}
	lastKey := v.keys[v.lastIndex]
	v.lastIndex = -1
	delete(v.Src.data, lastKey)
}

func (v *HashMapIterator[K, V]) Set(data KV[K, V]) KV[K, V] {
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
	if lastKey != data.Key() {
		panic("Map iterator.Set must set the same key!")
	}
	lastValue := v.Src.data[lastKey]
	delete(v.Src.data, lastKey)

	v.Src.data[data.Key()] = data.Value()
	v.keys[lastIndex] = data.Key()
	return KVOf(lastKey, lastValue)
}

func readKeys[K comparable, V any](mp map[K]V) []K {
	result := make([]K, len(mp))
	index := 0
	for key := range mp {
		result[index] = key
		index++
	}
	return result
}

func NewMapIteratorFor[K comparable, V any](v *HashMap[K, V]) coll.Iterator[KV[K, V]] {
	return &HashMapIterator[K, V]{Src: v, generation: v.generation, currentIndex: 0, lastIndex: -1, keys: readKeys(v.data)}
}

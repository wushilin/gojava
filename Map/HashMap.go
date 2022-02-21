package Map

import (
	"fmt"

	coll "github.com/wushilin/gojava/Collection"
	list "github.com/wushilin/gojava/List"
	set "github.com/wushilin/gojava/Set"
	"github.com/wushilin/stream"
)

// This interface represents a key value pair
type KV[K comparable, V any] interface {
	// Return the key of the entry
	Key() K

	// Return the value of the entry
	Value() V
}

type MapKV[K comparable, V any] struct {
	key   K
	value V
}

func (arg *MapKV[K, V]) Key() K {
	return arg.key
}

func (arg *MapKV[K, V]) Value() V {
	return arg.value
}

// Create a KV from KV value pair
func KVOf[K comparable, V any](key K, value V) KV[K, V] {
	return &MapKV[K, V]{key: key, value: value}
}

// This represents a Map interface
type Map[K comparable, V any] interface {
	// Return the count of objects
	Size() int

	// Test if a key is in the Map
	Contains(key K) bool

	// Get Value By Key, if no result found, ok is set to false
	Get(key K) (value V, ok bool)

	// Put Value By Key
	Put(key K, value V)

	// Put all
	PutAll(val Map[K, V])

	// Remove by key
	Remove(key K)

	// Remove all keys
	RemoveAll(keys coll.Collection[K])

	// Test if a value is found. This is convenient, but inefficient. It iterates the values. You can pass a Equalizer for teting object equality
	ContainsValueFunc(value V, equals coll.Equalizer[V]) bool

	// Test if a value is found. This is convenient, but inefficient. It iterates the values
	ContainsValue(value V) bool

	// Create iterator over a snapshot of the key set, but values are read in realtime. Calling Remove() deletes it from the Map
	// Calling Set(KV[K,V]) deletes the entry from map, but adds new kv to the map. Note that Set iterator's Set Function must set with same key.
	Iterator() coll.Iterator[KV[K, V]]

	// Return set of  Keys. It will not have duplicates. It is a snapshot, so calling remove for the set does nothing useful for you
	// Set has no order, and it does not guarantee the order is same as Values()
	Keys() set.Set[K]

	// Return collection of Values. Duplications might be there. It is a snapshot,  Calling Remove() or Set() does nothing useful for you.
	// The values has no order.
	Values() coll.Collection[V]

	// Return stream of KV[K,V]. It uses iterator internally
	Stream() stream.Stream[KV[K, V]]
}

type HashMap[K comparable, V any] struct {
	data       map[K]V
	generation int
}

func (v *HashMap[K, V]) applyMod() {
	v.generation++
}

func (v *HashMap[K, V]) Size() int {
	return len(v.data)
}

func (v *HashMap[K, V]) Contains(key K) bool {
	_, ok := v.data[key]
	return ok
}

func (v *HashMap[K, V]) Get(key K) (result V, ok bool) {
	result, ok = v.data[key]
	return
}

func (v *HashMap[K, V]) Put(key K, value V) {
	defer v.applyMod()
	v.data[key] = value
}

func (v *HashMap[K, V]) Clear() {
	defer v.applyMod()
	v.data = make(map[K]V)
}

func (v *HashMap[K, V]) ContainsValue(what V) bool {
	return v.ContainsValueFunc(what, coll.DefaultEqualizer[V]())
}

func (v *HashMap[K, V]) ContainsValueFunc(what V, equals coll.Equalizer[V]) bool {
	return v.Values().ContainsFunc(what, equals)
}

func (v *HashMap[K, V]) Values() coll.Collection[V] {
	result := list.NewLinkedList[V]()
	for _, v := range v.data {
		result.Add(v)
	}
	return result
}

func (v *HashMap[K, V]) PutAll(other Map[K, V]) {
	coll.ForEach(
		other.Iterator(),
		func(i KV[K, V]) bool {
			v.Put(i.Key(), i.Value())
			return true
		})
}
func (v *HashMap[K, V]) Remove(k K) {
	defer v.applyMod()
	delete(v.data, k)

}

func (v *HashMap[K, V]) RemoveAll(keys coll.Collection[K]) {
	coll.ForEach(keys.Iterator(), func(i K) bool {
		v.Remove(i)
		return true
	})
}
func (v *HashMap[K, V]) Keys() set.Set[K] {
	result := set.NewHashSet[K]()
	for k, _ := range v.data {
		result.Add(k)
	}
	return result
}

func (v *HashMap[K, V]) Iterator() coll.Iterator[KV[K, V]] {
	return NewMapIteratorFor(v)
}

func (v *HashMap[K, V]) Stream() stream.Stream[KV[K, V]] {
	return stream.FromIterator[KV[K, V]](v.Iterator())
}

func PrintMap[K comparable, V any](v Map[K, V]) {
	iter := v.Iterator()
	count := 0
	fmt.Print("{")
	for nextEntry, ok := iter.Next(); ok; nextEntry, ok = iter.Next() {
		key := nextEntry.Key()
		value := nextEntry.Value()
		format := "%+v=%+v,"
		if count == v.Size()-1 {
			format = format[:len(format)-1]
		}
		fmt.Print(fmt.Sprintf(format, key, value))
		count++
	}
	fmt.Println("}")
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &HashMap[K, V]{data: make(map[K]V), generation: 0}
}

# gojava
A few useful classes in Java implemented for Go

# String
```go
var hello String = String("Hello, 世界")
hello.Length() => 9 (9 unicode char)
string(hello) => convert to string

hello.ToCharArray() => []rune
hello.Concat(String("hey")) => "Hello, 世界hey"

hello.Contains("hey") => true
hello.EndsWith("hey") => true
hello.StartsWith("Hell1") => false
hello.SubStringWithLength(0, 4) => "Hell"
hello.SubString(9) => "hey"
String(",").Join([]{"hello", "world"}) => "hello,world"
hello.IndexOfString("hey") => 9
hello.LastIndexOf('e') => 10
hello.Matches("ello") => true
hello.ReplaceFirst("e", "eeeee") => "heeeeello, 世界hey"
hello.ReplaceAll("e", "f") => "Hfllo, 世界hfy"
hello.Repeat(2) => "Hello, 世界hey"
hello.Split("e") => ["H", "llo, 世界h", "y"]
hello.SplitLimit("e", 2) => ["H", "llo, 世界hey"]
```

# Collection
Interface for Collection
```go
type Iterator[T any] interface {
	// Iterator must be a stream.Iterator, which defines Next() T, bool
	Next() (element T, ok bool)

	// Remove last returned entry (can only be called once for every Next() call)
	Remove()

	// Set the value of current entry to the new vale, returns the old value
	Set(T) T
}

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
```
## List
NewLinkedList[T]()
LinkedListOf[T](args...T)
NewArrayList[T]()
ArrayListOf[T](arg...T)

```go
# Lists also supports other additional methods
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
```
They support all methods above, including iterator

## Set
```go
// Set is just a collection, but unlike List, it does not contain duplicates
type Set[T comparable] interface {
	coll.Collection[T]
}

// Set's Iterator.Set() removes current value, and adds a new Value. New value is not visible in current iteration.
// Multiple Set() on same spot replace the previously set value, or replace another value that will be iterated later.
// Use with care.
```
NewHashSet[T]()
HashSetOf[T](args...T)

It supports all methods above as defined by Collection.

# Map
```go
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
```


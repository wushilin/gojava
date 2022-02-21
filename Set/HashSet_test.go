package Set

import (
	"fmt"
	"testing"

	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/gojava/common"
	"github.com/wushilin/stream"
)

func TestHashset(t *testing.T) {
	set := NewHashSet[int]()
	stream.Range(1, 1000).Each(
		func(i int) {
			set.Add(i)
		})

	coll.PrintCollection[int](set)
	fmt.Println("Size: ", set.Size())

	set1 := HashSetOf(1, 2, 3, 4, 5)
	common.AssertTrue(t, set1.Contains(1))
	common.AssertFalse(t, set1.Contains(6))

	set2 := HashSetOf(3, 4, 5, 6, 7)

	set1Copy := NewHashSet[int]()
	set1Copy.AddAll(set1)

	set2Copy := NewHashSet[int]()
	set2Copy.AddAll(set2)

	coll.PrintCollection[int](set1)
	coll.PrintCollection[int](set2)

	set1.AddAll(set2)

	coll.PrintCollection[int](set1)
	set1.Clear()
	set1.AddAll(set1Copy)
	coll.PrintCollection[int](set1)

	set1.RemoveAll(set2)
	coll.PrintCollection[int](set1)
	set1.Clear()

	set1.AddAll(set1Copy)

	set1.RetainAll(set2)
	coll.PrintCollection[int](set1)

	join := NewHashSet[int]()
	join.AddAll(set1)

	set1.Clear()
	set1.AddAll(set1Copy)

	fmt.Println(set1.ContainsAll(join))
	set.Stream().Filter(func(i int) bool {
		return i%133 == 0
	}).Each(print)

	set4 := HashSetOf(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	iter4 := set4.Iterator()

	for item, ok := iter4.Next(); ok; item, ok = iter4.Next() {
		fmt.Println("Setting", item, "to", 5)
		//iter4.Remove()
		iter4.Set(5)
		coll.PrintCollection[int](set4)
	}
}

func print(i int) {
	fmt.Println(i)
}

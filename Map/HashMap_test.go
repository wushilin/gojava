package Map

import (
	"fmt"
	"testing"
	"time"

	coll "github.com/wushilin/gojava/Collection"
	"github.com/wushilin/gojava/common"
	"github.com/wushilin/stream"
)

func TestHashMap(t *testing.T) {
	mp := NewHashMap[int, int]()

	for i := 0; i < 100; i++ {
		mp.Put(i, 100-i)
	}
	PrintMap[int, int](mp)

	stream.Range(0, 100).Each(
		func(i int) {
			common.AssertTrue(t, mp.Contains(i))
		})
	stream.Range(0, 100).Each(
		func(i int) {
			value, ok := mp.Get(i)
			common.AssertTrue(t, ok)
			common.AssertEq(t, value, 100-i)
		})

	coll.PrintCollection[int](mp.Keys())
	coll.PrintCollection(mp.Values())

	fmt.Println("Before Size", mp.Size())
	iter := mp.Iterator()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		key := next.Key()
		value := next.Value()
		iter.Remove()
		iter.Set(KVOf(key, value+1000))
		iter.Set(KVOf(key, value+1001))
		iter.Set(KVOf(key, value+1002))
		iter.Set(KVOf(key, value+1003))
		iter.Set(KVOf(key, value+1004))
		fmt.Println("Found", key, "=", value)
	}

	count := 0
	mp.Stream().Each(func(val KV[int, int]) {
		fmt.Println("Visiting using stream ", val.Key(), "=>", val.Value())
		time.Sleep(100 * time.Millisecond)
		count++
	})
	fmt.Println("Visited", count, "entries")
	PrintMap[int, int](mp)
}

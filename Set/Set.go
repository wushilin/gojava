package Set

import coll "github.com/wushilin/gojava/Collection"

type Set[T comparable] interface {
	coll.Collection[T]
}

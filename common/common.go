package common

import (
	"log"
	"testing"
)

func AssertArrEq[T comparable](t *testing.T, a []T, b []T) {
	if len(a) != len(b) {
		t.Fatalf("Array length differ")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			log.Fatalf("Array differ")
		}
	}
}
func AssertTrue(t *testing.T, flag bool) {
	if !flag {
		t.Fatalf("Expect true but got false")
	}
}

func AssertFalse(t *testing.T, flag bool) {
	if flag {
		t.Fatalf("Expect false but got true")
	}
}
func AssertEq[T comparable](t *testing.T, a T, b T) {
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}

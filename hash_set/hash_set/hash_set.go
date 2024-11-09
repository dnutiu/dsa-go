package hash_set

import (
	"fmt"
	"strings"
)

type HashKey interface {
	~string | ~int | ~uint | ~int64 | ~uint64 | ~int32 | ~uint32 | ~int16 | ~uint16 | ~int8 | ~uint8
}

type Hashable[H HashKey] interface {
	Hash() H
}

// MyHashSet is a hash set implementation.
type MyHashSet[T Hashable[H], H HashKey] struct {
	storage map[H]T
}

// NewSet initializes a new hash set.
func NewSet[T Hashable[H], H HashKey]() MyHashSet[T, H] {
	return MyHashSet[T, H]{
		storage: make(map[H]T, 100),
	}
}

// Add adds an element to a set.
func (s *MyHashSet[T, H]) Add(element T) {
	// Hash element
	hash := element.Hash()

	// Save
	_, ok := s.storage[hash]
	if !ok {
		s.storage[hash] = element
	}
}

// AddAll adds all the elements to the set.
func (s *MyHashSet[T, H]) AddAll(elements ...T) {
	for _, element := range elements {
		s.Add(element)
	}
}

// Contains checks if the hash set contains the element T.
// Returns true if the element is part of the set, false otherwise.
func (s *MyHashSet[T, H]) Contains(element T) bool {
	// Hash element
	hash := element.Hash()
	_, ok := s.storage[hash]
	return ok
}

// Delete deletes an element from the set.
// Returns true if the element was deleted, false otherwise.
func (s *MyHashSet[T, H]) Delete(element T) bool {
	// Hash element
	hash := element.Hash()
	_, ok := s.storage[hash]
	if ok {
		delete(s.storage, hash)
		return true
	}
	return false
}

// String returns the string representation of the set.
func (s *MyHashSet[T, H]) String() string {
	var sb strings.Builder
	sb.WriteString("MyHashSet{")
	for _, value := range s.storage {
		sb.WriteString(fmt.Sprintf("%v,", value))
	}
	sb.WriteString("}")
	return sb.String()
}

// Union creates a union of two sets
func (s *MyHashSet[T, H]) Union(other *MyHashSet[T, H]) {
	for _, v := range other.storage {
		s.Add(v)
	}
}

// Sub creates a difference of two sets
func (s *MyHashSet[T, H]) Sub(other *MyHashSet[T, H]) {
	for _, v := range other.storage {
		s.Delete(v)
	}
}

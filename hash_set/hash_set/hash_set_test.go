package hash_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyString struct {
	Value string
}

func (m MyString) Hash() string {
	return m.Value
}

func TestNewSet(t *testing.T) {
	newSet := NewSet[MyString, string]()
	assert.NotNil(t, newSet)
}

func TestMyHashSet_Add(t *testing.T) {
	// Given
	newSet := NewSet[MyString, string]()
	newSet.Add(MyString{Value: "some"})

	// Then
	assert.True(t, newSet.Contains(MyString{Value: "some"}))
}

func TestMyHashSet_AddAll(t *testing.T) {
	// Given
	newSet := NewSet[MyString, string]()
	newSet.AddAll(MyString{Value: "some"}, MyString{Value: "another"})

	// Then
	assert.True(t, newSet.Contains(MyString{Value: "some"}))
	assert.True(t, newSet.Contains(MyString{Value: "another"}))
}

func TestMyHashSet_Delete(t *testing.T) {
	// Given
	newSet := NewSet[MyString, string]()
	newSet.AddAll(MyString{Value: "some"}, MyString{Value: "another"})
	newSet.Delete(MyString{Value: "some"})

	// Then
	assert.False(t, newSet.Contains(MyString{Value: "some"}))
	assert.True(t, newSet.Contains(MyString{Value: "another"}))
}

func TestMyHashSet_Union(t *testing.T) {
	// Setup
	newSet := NewSet[MyString, string]()
	newSet.AddAll(MyString{Value: "some"}, MyString{Value: "another"})

	anotherSet := NewSet[MyString, string]()
	anotherSet.AddAll(MyString{Value: "Batman"}, MyString{Value: "Robin"})

	// Test
	newSet.Union(&anotherSet)

	// Then
	assert.True(t, newSet.Contains(MyString{Value: "some"}))
	assert.True(t, newSet.Contains(MyString{Value: "another"}))
	assert.True(t, newSet.Contains(MyString{Value: "Batman"}))
	assert.True(t, newSet.Contains(MyString{Value: "Robin"}))
}

func TestMyHashSet_Sub(t *testing.T) {
	// Setup
	newSet := NewSet[MyString, string]()
	newSet.AddAll(MyString{Value: "some"}, MyString{Value: "another"}, MyString{Value: "Batman"})

	anotherSet := NewSet[MyString, string]()
	anotherSet.AddAll(MyString{Value: "Batman"}, MyString{Value: "Robin"})

	// Test
	newSet.Sub(&anotherSet)

	// Then
	assert.True(t, newSet.Contains(MyString{Value: "some"}))
	assert.True(t, newSet.Contains(MyString{Value: "another"}))
	assert.False(t, newSet.Contains(MyString{Value: "Batman"}))
	assert.False(t, newSet.Contains(MyString{Value: "Robin"}))
}

type String string

func (a String) Hash() string {
	return string(a)
}

func TestMyHashSet_Constraint(t *testing.T) {
	// Setup
	newSet := NewSet[String, string]()

	// Test
	newSet.AddAll("type", "constraints", "rock")

	// Then
	assert.True(t, newSet.Contains("type"))
	assert.True(t, newSet.Contains("constraints"))
	assert.True(t, newSet.Contains("rock"))
}

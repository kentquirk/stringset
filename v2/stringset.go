// Package stringset provides the capabilities of a standard Set with strings, including
// the standard set operations.
//
// The API is designed to be chainable so that common operations become one-liners,
// and it is designed to be easy to use with slices of strings.
//
// This package does not return errors. Set operations should be fast and
// chainable; adding items more than once, or attempting to remove things
// that don't exist are not errors.
package stringset

// this is an empty object that takes up no space
type nothing struct{}

// A StringSet is implemented using a map[string]nothing -- strings are copied
// when added to the set.
type StringSet struct {
	content map[string]nothing
}

// New constructs an empty StringSet.
func New() *StringSet {
	ss := new(StringSet)
	ss.content = make(map[string]nothing)
	return ss
}

// Add puts one or more strings into the set.
// If a string is already present the set is unchanged.
func (ss *StringSet) Add(sa ...string) *StringSet {
	for _, s := range sa {
		ss.content[s] = nothing{}
	}
	return ss
}

// Delete removes one or more items from a set if they exist in the set.
func (ss *StringSet) Delete(sa ...string) *StringSet {
	for _, s := range sa {
		delete(ss.content, s)
	}
	return ss
}

// Length returns the number of items currently in the set.
func (ss *StringSet) Length() int {
	return len(ss.content)
}

// Contains returns whether a given string is in the set.
func (ss *StringSet) Contains(s string) bool {
	_, ok := ss.content[s]
	return ok
}

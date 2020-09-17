package stringset

// Clone duplicates a StringSet.
func (ss *StringSet) Clone() *StringSet {
	clone := New()
	for k := range ss.content {
		clone.content[k] = ss.content[k]
	}
	return clone
}

// Intersection returns the symmetric difference of the two sets.
// This returns a new set -- both operands are left unchanged.
//
// abc & cde == c
func (ss *StringSet) Intersection(ss2 *StringSet) *StringSet {
	// we can optimize for length by iterating on the smaller set
	l1 := len(ss.content)
	l2 := len(ss2.content)
	if l2 < l1 {
		ss2, ss = ss, ss2
	}

	intersection := New()
	for k := range ss.content {
		if _, ok := ss2.content[k]; ok {
			intersection.Add(k)
		}
	}
	return intersection
}

// Union generates the union (OR) of the two sets.
// Returns a new set and leaves the operands unchanged.
func (ss *StringSet) Union(ss2 *StringSet) *StringSet {
	union := New()
	for k := range ss.content {
		union.Add(k)
	}
	for k := range ss2.content {
		union.Add(k)
	}
	return union
}

// Difference is an assymetric set difference.
// It subtracts the rhs from the lhs and returns a new set.
func (ss *StringSet) Difference(ss2 *StringSet) *StringSet {
	difference := New()
	for k := range ss.content {
		if _, ok := ss2.content[k]; !ok {
			difference.Add(k)
		}
	}
	return difference
}

// Equals checks whether two string sets have the same members.
func (ss *StringSet) Equals(ss2 *StringSet) bool {
	if ss.Length() != ss2.Length() {
		return false
	}
	// if two sets have the same length then we know that if there is a difference,
	// it must be the case that there are the same number of unique members
	// in each set. So we can iterate over one set, and know that if every
	// item was found, the two sets are equivalent
	for k := range ss.content {
		if _, ok := ss2.content[k]; !ok {
			return false
		}
	}
	return true
}

// Filter returns a new StringSet that contains only the members of
// the original set where f(member) evaluates to true.
func (ss *StringSet) Filter(f func(s string) bool) *StringSet {
	filtered := New()
	for k := range ss.content {
		if ok := f(k); ok {
			filtered.Add(k)
		}
	}
	return filtered
}

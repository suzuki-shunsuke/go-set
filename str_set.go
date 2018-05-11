package set

import (
	"encoding/json"
)

// StrSet represents a string set.
type StrSet struct {
	data map[string]struct{}
}

// NewStrSet returns a StrSet.
func NewStrSet(args ...string) *StrSet {
	set := &StrSet{data: map[string]struct{}{}}
	for _, a := range args {
		set.data[a] = struct{}{}
	}
	return set
}

// Len returns the number of elements of the set.
// If set is nil, 0 is returned.
func (set *StrSet) Len() int {
	if set == nil {
		return 0
	}
	return len(set.data)
}

// Has returns whether a given string is included in the set.
// If set is nil, false is returned.
func (set *StrSet) Has(k string) bool {
	if set == nil {
		return false
	}
	if set.data == nil {
		return false
	}
	_, ok := set.data[k]
	return ok
}

// HasAll returns whether all given strings are included in the set.
// If set is nil, false is returned.
func (set *StrSet) HasAll(args ...string) bool {
	for _, a := range args {
		if !set.Has(a) {
			return false
		}
	}
	return true
}

// HasAny returns whether some given strings are included in the set.
// If set is nil, false is returned.
func (set *StrSet) HasAny(args ...string) bool {
	for _, a := range args {
		if set.Has(a) {
			return true
		}
	}
	return false
}

// Add adds a string to the set.
func (set *StrSet) Add(k string) {
	if set.data == nil {
		set.data = map[string]struct{}{}
	}
	set.data[k] = struct{}{}
}

// Adds adds strings to the set.
func (set *StrSet) Adds(args ...string) {
	if set.data == nil {
		set.data = map[string]struct{}{}
	}
	for _, k := range args {
		set.data[k] = struct{}{}
	}
}

// AddSet adds a StrSet to the set.
func (set *StrSet) AddSet(other *StrSet) {
	if other == nil {
		return
	}
	if set.data == nil {
		set.data = map[string]struct{}{}
	}
	for k := range other.data {
		set.data[k] = struct{}{}
	}
}

// AddSets adds StrSets to the set.
func (set *StrSet) AddSets(others ...*StrSet) {
	if set.data == nil {
		set.data = map[string]struct{}{}
	}
	for _, other := range others {
		if other == nil {
			continue
		}
		for k := range other.data {
			set.data[k] = struct{}{}
		}
	}
}

// Clone returns a new StrSet which has same elements.
// If set is nil, new empty set is returned.
func (set *StrSet) Clone() *StrSet {
	if set == nil {
		return NewStrSet()
	}
	if set.data == nil {
		return &StrSet{}
	}
	s := NewStrSet()
	for k := range set.data {
		s.data[k] = struct{}{}
	}
	return s
}

// Remove removes a string from the set.
// If set is nil, nothing happens.
func (set *StrSet) Remove(k string) {
	if set == nil {
		return
	}
	delete(set.data, k)
}

// Removes removes strings from the set.
// If set is nil, nothing happens.
func (set *StrSet) Removes(args ...string) {
	if set == nil {
		return
	}
	for _, k := range args {
		delete(set.data, k)
	}
}

// Clear removes all elements.
// If set is nil, nothing happens.
func (set *StrSet) Clear() {
	if set == nil {
		return
	}
	set.data = map[string]struct{}{}
}

// MarshalJSON is the implementation of the json.Marshaler interface.
func (set *StrSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToList())
}

// UnmarshalJSON is the implementation of the json.Unmarshaler interface.
func (set *StrSet) UnmarshalJSON(b []byte) error {
	arr := []string{}
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return err
	}
	if set.data == nil {
		set.data = map[string]struct{}{}
	}
	for _, k := range arr {
		set.data[k] = struct{}{}
	}
	return nil
}

// ToList returns a list composed of elements of the set.
func (set *StrSet) ToList() []string {
	if set == nil {
		return nil
	}
	size := len(set.data)
	if size == 0 {
		return nil
	}
	arr := make([]string, size)
	i := 0
	for k := range set.data {
		arr[i] = k
		i++
	}
	return arr
}

// ToMap returns a map whose keys are elements of the set.
// If the parameter 'deep' is true, this method returns a new map.
// If the parameter 'deep' is false, this method returns the map which the set has internally.
func (set *StrSet) ToMap(deep bool) map[string]struct{} {
	if set == nil {
		return nil
	}
	if !deep {
		return set.data
	}
	m := map[string]struct{}{}
	for k := range set.data {
		m[k] = struct{}{}
	}
	return m
}

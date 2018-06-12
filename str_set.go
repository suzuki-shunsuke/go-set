package set

import (
	"encoding/json"
	"fmt"
)

// StrSet represents a string set.
type StrSet map[string]struct{}

// NewStrSet returns a StrSet.
func NewStrSet(args ...string) StrSet {
	set := StrSet{}
	for _, a := range args {
		set[a] = struct{}{}
	}
	return set
}

// Has returns whether a given string is included in the set.
// If set is nil, false is returned.
func (set StrSet) Has(k string) bool {
	_, ok := set[k]
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
// An error is returned when set is nil.
func (set StrSet) Add(k string) error {
	if set == nil {
		return fmt.Errorf("set is nil")
	}
	set[k] = struct{}{}
	return nil
}

// Adds adds strings to the set.
// An error is returned when set is nil.
func (set StrSet) Adds(args ...string) error {
	if len(args) == 0 {
		return nil
	}
	if set == nil {
		return fmt.Errorf("set is nil")
	}
	for _, k := range args {
		set[k] = struct{}{}
	}
	return nil
}

// AddSet adds a StrSet to the set.
// An error is returned when set is nil.
func (set StrSet) AddSet(other StrSet) error {
	if len(other) == 0 {
		return nil
	}
	if set == nil {
		return fmt.Errorf("set is nil")
	}
	for k := range other {
		set[k] = struct{}{}
	}
	return nil
}

// AddSets adds StrSets to the set.
// An error is returned when set is nil.
func (set StrSet) AddSets(others ...StrSet) error {
	if len(others) == 0 {
		return nil
	}
	if set == nil {
		return fmt.Errorf("set is nil")
	}
	for _, other := range others {
		if other == nil {
			continue
		}
		for k := range other {
			set[k] = struct{}{}
		}
	}
	return nil
}

// Clone returns a new StrSet which has same elements.
// If set is nil, new empty set is returned.
func (set StrSet) Clone() StrSet {
	s := StrSet{}
	for k := range set {
		s[k] = struct{}{}
	}
	return s
}

// Remove removes a string from the set.
// If set is nil, nothing happens.
func (set StrSet) Remove(k string) {
	if set == nil {
		return
	}
	delete(set, k)
}

// Removes removes strings from the set.
// If set is nil, nothing happens.
func (set StrSet) Removes(args ...string) {
	if set == nil {
		return
	}
	for _, k := range args {
		delete(set, k)
	}
}

// Clear removes all elements.
// If set is nil, nothing happens.
func (set StrSet) Clear() {
	if set == nil {
		return
	}
	for k := range set {
		delete(set, k)
	}
}

// MarshalJSON is the implementation of the json.Marshaler interface.
func (set StrSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToList())
}

// UnmarshalJSON is the implementation of the json.Unmarshaler interface.
// An error is returned when set is nil.
func (set StrSet) UnmarshalJSON(b []byte) error {
	arr := []string{}
	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}
	if len(arr) == 0 {
		return nil
	}
	if set == nil {
		return fmt.Errorf("set is nil")
	}
	for _, k := range arr {
		set[k] = struct{}{}
	}
	return nil
}

// ToList returns a list composed of elements of the set.
// If set is nil, an empty list is returned.
func (set StrSet) ToList() []string {
	size := len(set)
	if size == 0 {
		return []string{}
	}
	arr := make([]string, size)
	i := 0
	for k := range set {
		arr[i] = k
		i++
	}
	return arr
}

// ToMap returns a map whose keys are elements of the set.
// If the parameter 'deep' is true, this method returns a new map.
// If the parameter 'deep' is false, this method returns the map which the set has internally.
// If set is nil, an empty map is returned.
func (set StrSet) ToMap(deep bool) map[string]struct{} {
	if set == nil {
		return map[string]struct{}{}
	}
	if !deep {
		return map[string]struct{}(set)
	}
	m := map[string]struct{}{}
	for k := range set {
		m[k] = struct{}{}
	}
	return m
}

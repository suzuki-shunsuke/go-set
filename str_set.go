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
func (set *StrSet) Len() int {
	return len(set.data)
}

// Has returns whether a given string is included in the set.
func (set *StrSet) Has(k string) bool {
	if set.data == nil {
		return false
	}
	_, ok := set.data[k]
	return ok
}

// HasAll returns whether all given strings are included in the set.
func (set *StrSet) HasAll(args ...string) bool {
	for _, a := range args {
		if !set.Has(a) {
			return false
		}
	}
	return true
}

// Has returns whether some given strings are included in the set.
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

// Remove removes a string from the set.
func (set *StrSet) Remove(k string) {
	delete(set.data, k)
}

// Removes removes strings from the set.
func (set *StrSet) Removes(args ...string) {
	for _, k := range args {
		delete(set.data, k)
	}
}

// Clear removes all elements.
func (set *StrSet) Clear() {
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
	arr := make([]string, len(set.data))
	i := 0
	for k, _ := range set.data {
		arr[i] = k
		i++
	}
	return arr
}

// ToMap returns a map whose keys are elements of the set.
// If the parameter 'deep' is true, this method returns a new map.
// If the parameter 'deep' is false, this method returns the map which the set has internally.
func (set *StrSet) ToMap(deep bool) map[string]struct{} {
	if !deep {
		return set.data
	}
	m := map[string]struct{}{}
	for k, _ := range set.data {
		m[k] = struct{}{}
	}
	return m
}

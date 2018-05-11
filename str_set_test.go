package set_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-set"
)

func TestStrSetLen(t *testing.T) {
	s := set.NewStrSet("foo", "bar")
	if s.Len() != 2 {
		t.Fatalf(`s.Len() = %d, wanted 2`, s.Len())
	}
	s = nil
	if s.Len() != 0 {
		t.Fatalf(`s.Len() = %d, wanted 0`, s.Len())
	}
}

func TestStrSetHas(t *testing.T) {
	s := set.NewStrSet()
	k := "hello"
	if s.Has(k) {
		t.Fatal(s)
	}
	s.Add(k)
	if !s.Has(k) {
		t.Fatal(s)
	}
	s = &set.StrSet{}
	if s.Has(k) {
		t.Fatal(s)
	}
	s = nil
	if s.Has(k) {
		t.Fatal(s)
	}
}

func TestStrSetHasAll(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	if !s.HasAll("hello") {
		t.Fatal(s)
	}
	if !s.HasAll("foo") {
		t.Fatal(s)
	}
	if !s.HasAll("foo", "hello") {
		t.Fatal(s)
	}
	if s.HasAll("foo", "bar") {
		t.Fatal(s)
	}
}

func TestStrSetHasAny(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	if !s.HasAny("hello") {
		t.Fatal(s)
	}
	if !s.HasAny("foo") {
		t.Fatal(s)
	}
	if !s.HasAny("foo", "hello") {
		t.Fatal(s)
	}
	if !s.HasAny("foo", "bar") {
		t.Fatal(s)
	}
	if s.HasAny("bar") {
		t.Fatal(s)
	}
}

func TestStrSetAdd(t *testing.T) {
	s := &set.StrSet{}
	if err := s.Add("hello"); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s = nil
	if err := s.Add("hello"); err == nil {
		t.Fatal("set is nil: ", s)
	}
}

func TestStrSetAdds(t *testing.T) {
	s := set.NewStrSet("hello", "zoo")
	if err := s.Adds("hello", "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 4 {
		t.Fatal(s)
	}
	s = &set.StrSet{}
	if err := s.Adds("hello", "foo"); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 2 {
		t.Fatal(s)
	}
	s = nil
	if err := s.Adds("hello", "foo"); err == nil {
		t.Fatal("set is nil: ", s)
	}
}

func TestStrSetAddSet(t *testing.T) {
	s := &set.StrSet{}
	s2 := set.NewStrSet("a", "b")
	if err := s.AddSet(s2); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 2 {
		t.Fatal(s)
	}
	if s2.Len() != 2 {
		t.Fatal(s2)
	}
	var s3 *set.StrSet
	if err := s.AddSet(s3); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 2 {
		t.Fatal(s)
	}
	if s3 != nil {
		t.Fatal(s3)
	}
	s = nil
	if err := s.AddSet(s3); err == nil {
		t.Fatal("set is nil: ", s)
	}
}

func TestStrSetAddSets(t *testing.T) {
	s := &set.StrSet{}
	s2 := set.NewStrSet("a", "b")
	var s3 *set.StrSet
	if err := s.AddSets(s2, s3); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 2 {
		t.Fatal(s)
	}
	if s2.Len() != 2 {
		t.Fatal(s2)
	}
	if s3 != nil {
		t.Fatal(s3)
	}
	s = nil
	if err := s.AddSets(s2); err == nil {
		t.Fatal("set is nil: ", s)
	}
}

func TestStrSetClone(t *testing.T) {
	s := &set.StrSet{}
	s2 := s.Clone()
	if s2.Len() != 0 {
		t.Fatal(s2)
	}
	s.Add("foo")
	s2 = s.Clone()
	if s2.Len() != 1 {
		t.Fatal(s2)
	}
	s = nil
	s2 = s.Clone()
	if s2.Len() != 0 {
		t.Fatal(s2)
	}
}

func TestStrSetRemove(t *testing.T) {
	s := set.NewStrSet("hello")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s.Remove("foo")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s.Remove("hello")
	if s.Len() != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Remove("hello")
	if s.Len() != 0 {
		t.Fatal(s)
	}
}

func TestStrSetRemoves(t *testing.T) {
	s := set.NewStrSet("hello")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s.Removes("foo", "bar")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s.Removes("hello", "foo")
	if s.Len() != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Removes("hello", "foo")
	if s.Len() != 0 {
		t.Fatal(s)
	}
}

func TestStrSetClear(t *testing.T) {
	s := set.NewStrSet("hello", "bar")
	if s.Len() != 2 {
		t.Fatal(s)
	}
	s.Clear()
	if s.Len() != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Clear()
	if s.Len() != 0 {
		t.Fatal(s)
	}
}

func TestStrSetMarshalJSON(t *testing.T) {
	s := set.NewStrSet("hello")
	b, err := s.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `["hello"]` {
		t.Fatal(string(b))
	}
	s = nil
	b, err = s.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "null" {
		t.Fatal(string(b))
	}
}

func TestStrSetUnmarshalJSON(t *testing.T) {
	s := set.NewStrSet()
	if err := json.Unmarshal([]byte(`["foo", "bar", "foo"]`), s); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 2 {
		t.Fatal(s)
	}
	if err := json.Unmarshal([]byte(`"foo"`), s); err == nil {
		t.Fatal("Unmarshal should fail")
	}
	s = &set.StrSet{}
	if err := json.Unmarshal([]byte(`["foo", "bar", "foo"]`), s); err != nil {
		t.Fatal(err)
	}
}

func TestStrSetToList(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	arr := s.ToList()
	if len(arr) != 2 {
		t.Fatal(arr)
	}
	if size := len(set.NewStrSet().ToList()); size != 0 {
		t.Fatalf(`len(set.NewStrSet().ToList()) = %d, wanted 0`, size)
	}

	s = nil
	arr = s.ToList()
	if len(arr) != 0 {
		t.Fatal(arr)
	}
}

func TestStrSetToMap(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	m := s.ToMap(false)
	if len(m) != 2 {
		t.Fatal(m)
	}
	delete(m, "hello")
	if s.Len() != 1 {
		t.Fatal(s)
	}
	s = set.NewStrSet("hello", "foo")
	m = s.ToMap(true)
	delete(m, "hello")
	if s.Len() != 2 {
		t.Fatal(s)
	}

	s = nil
	m = s.ToMap(false)
	if len(m) != 0 {
		t.Fatal(m)
	}
	m = s.ToMap(true)
	if len(m) != 0 {
		t.Fatal(m)
	}
}

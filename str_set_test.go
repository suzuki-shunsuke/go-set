package set_test

import (
	"encoding/json"
	"testing"

	"github.com/suzuki-shunsuke/go-set"
)

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
}

func TestStrSetAdds(t *testing.T) {
	s := set.NewStrSet("hello", "zoo")
	if s.Len() != 2 {
		t.Fatal(s)
	}
	s.Adds("hello", "foo", "bar")
	if s.Len() != 4 {
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
}

func TestStrSetLen(t *testing.T) {
	s := set.NewStrSet("foo", "bar")
	if s.Len() != 2 {
		t.Fatalf(`s.Len() = %d, wanted 2`, s.Len())
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

func TestStrSetToList(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	arr := s.ToList()
	if len(arr) != 2 {
		t.Fatal(arr)
	}
}

func TestStrSetToMap(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	if s.Len() != 2 {
		t.Fatal(s)
	}
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
}

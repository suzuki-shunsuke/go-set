package set_test

import (
	"encoding/json"
	"reflect"
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
	s = set.StrSet{}
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
	s := set.StrSet{}
	if err := s.Add("hello"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.ToList(), []string{"hello"}) {
		t.Fatal(s)
	}
	s = nil
	if err := s.Add("hello"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.ToList(), []string{"hello"}) {
		t.Fatal(s)
	}
	var s2 *set.StrSet
	if err := s2.Add("hello"); err == nil {
		t.Fatal("set is nil: ", s2)
	}
}

func TestStrSetAdds(t *testing.T) {
	s := set.NewStrSet("hello", "zoo")
	if err := s.Adds("hello", "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	exp := map[string]struct{}{
		"hello": {},
		"zoo":   {},
		"foo":   {},
		"bar":   {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if len(s) != 4 {
		t.Fatal(s)
	}
	s = set.StrSet{}
	if err := s.Adds("hello", "foo"); err != nil {
		t.Fatal(err)
	}
	exp = map[string]struct{}{
		"hello": {},
		"foo":   {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	s = nil
	if err := s.Adds("hello", "foo"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if err := s.Adds(); err != nil {
		t.Fatal(err)
	}
	var s2 *set.StrSet
	if err := s2.Adds("hello", "foo"); err == nil {
		t.Fatal("set is nil", s2)
	}
}

func TestStrSetAddSet(t *testing.T) {
	s := set.StrSet{}
	s2 := set.NewStrSet("a", "b")
	if err := s.AddSet(s2); err != nil {
		t.Fatal(err)
	}
	exp := map[string]struct{}{
		"a": {},
		"b": {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if !reflect.DeepEqual(s2.ToMap(false), exp) {
		t.Fatal(s2)
	}
	var s3 set.StrSet
	if err := s.AddSet(s3); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if s3 != nil {
		t.Fatal(s3)
	}
	s = nil
	if err := s.AddSet(s3); err != nil {
		t.Fatal(err)
	}
	s3 = set.NewStrSet("a")
	if err := s.AddSet(s3); err != nil {
		t.Fatal(err)
	}
	var s4 *set.StrSet
	if err := s4.AddSet(s3); err == nil {
		t.Fatal("set is nil", s4)
	}
}

func TestStrSetAddSets(t *testing.T) {
	s := set.StrSet{}
	s2 := set.NewStrSet("a", "b")
	var s3 set.StrSet
	if err := s.AddSets(s2, s3); err != nil {
		t.Fatal(err)
	}
	exp := map[string]struct{}{
		"a": {},
		"b": {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if !reflect.DeepEqual(s2.ToMap(false), exp) {
		t.Fatal(s2)
	}
	if s3 != nil {
		t.Fatal(s3)
	}
	s = nil
	if err := s.AddSets(s2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s2.ToMap(false), exp) {
		t.Fatal(s2)
	}
	if err := s.AddSets(); err != nil {
		t.Fatal(err)
	}
	var s4 *set.StrSet
	if err := s4.AddSets(s2); err == nil {
		t.Fatal("set is nil", s4)
	}
}

func TestStrSetClone(t *testing.T) {
	s := set.StrSet{}
	s2 := s.Clone()
	if len(s2) != 0 {
		t.Fatal(s2)
	}
	s.Add("foo")
	s2 = s.Clone()
	exp := []string{"foo"}
	if !reflect.DeepEqual(s2.ToList(), exp) {
		t.Fatal(s2)
	}
	s = nil
	s2 = s.Clone()
	if len(s2) != 0 {
		t.Fatal(s2)
	}
}

func TestStrSetRemove(t *testing.T) {
	s := set.NewStrSet("hello")
	s.Remove("foo")
	exp := []string{"hello"}
	if !reflect.DeepEqual(s.ToList(), exp) {
		t.Fatal(s)
	}
	s.Remove("hello")
	if len(s) != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Remove("hello")
	if len(s) != 0 {
		t.Fatal(s)
	}
}

func TestStrSetRemoves(t *testing.T) {
	s := set.NewStrSet("hello")
	s.Removes("foo", "bar")
	exp := []string{"hello"}
	if !reflect.DeepEqual(s.ToList(), exp) {
		t.Fatal(s)
	}
	s.Removes("hello", "foo")
	if len(s) != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Removes("hello", "foo")
	if len(s) != 0 {
		t.Fatal(s)
	}
}

func TestStrSetClear(t *testing.T) {
	s := set.NewStrSet("hello", "bar")
	exp := map[string]struct{}{
		"hello": {},
		"bar":   {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	s.Clear()
	if len(s) != 0 {
		t.Fatal(s)
	}
	s = nil
	s.Clear()
	if len(s) != 0 {
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
	if string(b) != "[]" {
		t.Fatal(string(b))
	}
}

func TestStrSetUnmarshalJSON(t *testing.T) {
	s := set.NewStrSet()
	if err := json.Unmarshal([]byte(`["foo", "bar", "foo"]`), &s); err != nil {
		t.Fatal(err)
	}
	exp := map[string]struct{}{
		"foo": {},
		"bar": {},
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	if err := json.Unmarshal([]byte(`"foo"`), &s); err == nil {
		t.Fatal("Unmarshal should fail")
	}
	s = set.StrSet{}
	if err := json.Unmarshal([]byte(`["foo", "bar", "foo"]`), &s); err != nil {
		t.Fatal(err)
	}
	s = nil
	if err := json.Unmarshal([]byte(`["foo", "bar", "foo"]`), &s); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.ToMap(false), exp) {
		t.Fatal(s)
	}
	s = set.StrSet{}
	if err := json.Unmarshal([]byte(`[]`), &s); err != nil {
		t.Fatal(err)
	}
	a := struct {
		Set set.StrSet `json:"set"`
	}{}
	if err := json.Unmarshal([]byte(`{"set": ["foo", "bar", "foo"]}`), &a); err != nil {
		t.Fatal(err)
	}
	var s2 *set.StrSet = nil
	if err := s2.UnmarshalJSON([]byte(`["foo", "bar", "foo"]`)); err == nil {
		t.Fatal("set is nil", s2)
	}
}

func TestStrSetToList(t *testing.T) {
	s := set.NewStrSet("hello", "foo")
	arr := s.ToList()
	exp1 := []string{"hello", "foo"}
	exp2 := []string{"foo", "hello"}
	if !reflect.DeepEqual(arr, exp1) && !reflect.DeepEqual(arr, exp2) {
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
	exp := map[string]struct{}{
		"foo":   {},
		"hello": {},
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatal(m)
	}
	delete(m, "hello")
	if !reflect.DeepEqual(s.ToList(), []string{"foo"}) {
		t.Fatal(s)
	}
	s = set.NewStrSet("hello", "foo")
	m = s.ToMap(true)
	delete(m, "hello")
	if !reflect.DeepEqual(s.ToMap(false), exp) {
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

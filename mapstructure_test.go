package set_test

import (
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/suzuki-shunsuke/go-set"
)

func TestMapstructureDecodeHookFromListToStrSet(t *testing.T) {
	s := set.NewStrSet()
	input := []interface{}{"foo"}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: set.MapstructureDecodeHookFromListToStrSet,
		Result:     s,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := decoder.Decode(input); err != nil {
		t.Fatal(err)
	}
	if s.Len() != 1 {
		t.Fatal(s)
	}

	output := []string{}
	decoder2, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: set.MapstructureDecodeHookFromListToStrSet,
		Result:     &output,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := decoder2.Decode(input); err != nil {
		t.Fatal(err)
	}

	input2 := map[string]string{}
	if err := decoder.Decode(input2); err != nil {
		t.Fatal(err)
	}

	input3 := []interface{}{1}
	if err := decoder.Decode(input3); err == nil {
		t.Fatal(err)
	}
}

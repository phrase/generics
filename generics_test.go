package generics

import (
	"fmt"
	"sort"
	"testing"
)

type record struct {
	Name   string
	Amount int
}

func TestKeys(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	keys := Keys(m).([]string)
	sort.Strings(keys)
	tests := []struct{ Has, Want interface{} }{
		{len(keys), 3},
		{keys[0], "one"},
	}

	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}

}

func TestIndex(t *testing.T) {
	records := []*record{
		{Name: "one"},
		{Name: "two"},
		{Name: "three"},
	}

	m := Index(records, func(r *record) string { return r.Name }).(map[string]*record)

	tests := []struct{ Has, Want interface{} }{
		{len(m), 3},
		{m["one"].Name, "one"},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestReject(t *testing.T) {
	records := []*record{
		{Amount: 1},
		{Amount: 2},
		{Amount: 3},
	}
	filter := func(r *record) bool {
		return r.Amount >= 2
	}

	filtered := Reject(records, filter).([]*record)

	tests := []struct{ Has, Want interface{} }{
		{len(filtered), 1},
		{filtered[0].Amount, 1},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestSelect(t *testing.T) {
	records := []*record{
		{Amount: 1},
		{Amount: 2},
		{Amount: 3},
	}

	filter := func(r *record) bool {
		return r.Amount >= 2
	}

	filtered := Select(records, filter).([]*record)

	tests := []struct{ Has, Want interface{} }{
		{len(filtered), 2},
		{filtered[0].Amount, 2},
		{filtered[1].Amount, 3},
		{len(Select(records, func(r *record) bool { return false }).([]*record)), 0},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestTail(t *testing.T) {
	in := []string{"one", "two", "three"}
	records := []*record{{Name: "foo"}}

	tests := []struct{ Input, Want interface{} }{
		{Tail(in, 1).([]string), "[three]"},
		{Tail(in, 0).([]string), "[]"},
		{Tail(in, 2).([]string), "[two three]"},
		{Tail(in, 5).([]string), "[one two three]"},
		{Tail([]string{}, 5).([]string), "[]"},
		{Tail([]string{}, -1).([]string), "[]"},
		{Tail(records, 1).([]*record)[0].Name, "foo"},
	}
	for i, tc := range tests {
		has := fmt.Sprintf("%v", tc.Input)
		if tc.Want != has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, has, has)
		}
	}
}

func TestFirst(t *testing.T) {
	in := []*record{
		{Name: "one"},
		{Name: "two"},
		{Name: "three"},
	}
	r := First(in).(*record)

	tests := []struct{ Has, Want interface{} }{
		{r.Name, "one"},
		{First([]*record{}) == nil, true},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %q, was %q", i+1, tc.Want, tc.Has)
		}
	}
}

func TestLast(t *testing.T) {
	in := []*record{
		{Name: "one"},
		{Name: "two"},
		{Name: "three"},
	}
	r := Last(in).(*record)

	tests := []struct{ Has, Want interface{} }{
		{r.Name, "three"},
		{First([]*record{}) == nil, true},
		{First([]string{"first"}).(string), "first"},
		{First([]string{}).(string), ""},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %q, was %q", i+1, tc.Want, tc.Has)
		}
	}
}
func TestHead(t *testing.T) {
	in := []string{"one", "two", "three"}
	tests := []struct{ Input, Want interface{} }{
		{Head(in, 1).([]string), "[one]"},
		{Head(in, 0).([]string), "[]"},
		{Head(in, 2).([]string), "[one two]"},
		{Head(in, 5).([]string), "[one two three]"},
		{Head([]string{}, 5).([]string), "[]"},
		{Head([]string{}, -1).([]string), "[]"},
	}
	for i, tc := range tests {
		has := fmt.Sprintf("%v", tc.Input)
		if tc.Want != has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, has, has)
		}
	}
}

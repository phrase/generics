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

func TestSort(t *testing.T) {
	records := []*record{
		{Name: "one", Amount: 1},
		{Name: "two", Amount: 1},
		{Name: "three", Amount: 2},
	}

	t.Run("with attribute name", func(t *testing.T) {
		SortReverse(records, "Amount")
		if has := records[0].Name; has != "three" {
			t.Errorf("has was %q", has)
		}
	})

	t.Run("with function", func(t *testing.T) {
		SortReverse(records, func(r *record) int { return r.Amount })
		if has := records[0].Name; has != "three" {
			t.Errorf("has was %q", has)
		}
	})
}

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3})
	if sum != 6 {
		t.Errorf("sum was %v", sum)
	}
}

func TestMax(t *testing.T) {
	max := Max([]int{1, 2, 100, 3})
	if max != 100.0 {
		t.Errorf("max was %v", max)
	}
}

func TestToFloats(t *testing.T) {
	res := toFloats([]int{1, 2, 3, 4})
	if has := fmt.Sprintf("%#v", res); has != "[]float64{1, 2, 3, 4}" {
		t.Errorf("%s", has)
	}
	res = toFloats([]interface{}{1, 2.0, 3, 4.1})
	if has := fmt.Sprintf("%#v", res); has != "[]float64{1, 2, 3, 4.1}" {
		t.Errorf("%s", has)
	}
}

func TestFoldLeft(t *testing.T) {
	ints := []int{1, 2, 3}

	t.Run("with struct", func(t *testing.T) {
		type stat struct {
			Sum int
		}

		folder := func(s *stat, value int) *stat {
			s.Sum += value
			return s
		}
		res := FoldLeft(ints, folder).(*stat)
		if v := res.Sum; v != 6 {
			t.Errorf("sum is %d", v)
		}
	})

	t.Run("square of elements", func(t *testing.T) {
		folder := func(list []int, value int) []int { return append(list, value*value) }
		res := FoldLeft(ints, folder).([]int)
		if v := fmt.Sprintf("%v", res); v != "[1 4 9]" {
			t.Errorf("was %s", v)
		}
	})

	t.Run("sum", func(t *testing.T) {
		folder := func(sum int, value int) int { return sum + value }
		res := FoldLeft(ints, folder).(int)
		if res != 6 {
			t.Errorf("was %d", res)
		}
	})

	t.Run("with a map", func(t *testing.T) {
		type record struct {
			Name  string
			Count int
		}
		list := []*record{
			{"1", 1},
			{"2", 2},
			{"3", 3},
		}
		folder := func(m map[string]int, r *record) map[string]int {
			m[r.Name] = r.Count * 2
			return m
		}
		res := FoldLeft(list, folder).(map[string]int)
		if v := res["1"]; v != 2 {
			t.Errorf("%d", v)
		}
		if v := res["2"]; v != 4 {
			t.Errorf("%d", v)
		}
	})
}

func TestGroup(t *testing.T) {
	records := []*record{
		{Name: "one", Amount: 1},
		{Name: "two", Amount: 1},
		{Name: "three", Amount: 2},
	}
	res := Group(records, "Amount").(map[int][]*record)
	tests := []struct{ Has, Want interface{} }{
		{len(res), 2},
		{len(res[1]), 2},
		{res[1][0].Name, "one"},
		{res[1][1].Name, "two"},
		{res[2][0].Name, "three"},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestMap(t *testing.T) {
	records := []*record{
		{Name: "one", Amount: 1},
		{Name: "two", Amount: 1},
		{Name: "three", Amount: 2},
	}
	amounts := Map(records, "Amount").([]int)
	mapper := func(r *record) string { return r.Name }
	names := Map(records, mapper).([]string)

	tests := []struct{ Has, Want interface{} }{
		{fmt.Sprintf("%+v", names), "[one two three]"},
		{len(amounts), 3},
		{amounts[0], 1},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestAttribute(t *testing.T) {
	records := []*record{
		{Name: "Parrot", Amount: 3},
		{Name: "PhraseApp", Amount: 2},
	}

	names := Attributes(records, "Name").([]string)
	amounts := Attributes(records, "Amount").([]int)

	tests := []struct{ Has, Want interface{} }{
		{len(names), 2},
		{names[0], "Parrot"},
		{len(amounts), 2},
		{amounts[0], 3},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
}

func TestValues(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	values := Values(m).([]string)
	sort.Strings(values)

	tests := []struct{ Has, Want interface{} }{
		{len(values), 3},
		{values[0], "one"},
	}
	for i, tc := range tests {
		if tc.Want != tc.Has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, tc.Has, tc.Has)
		}
	}
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
	m2 := Index(records, "Name").(map[string]*record)

	tests := []struct{ Has, Want interface{} }{
		{len(m), 3},
		{m["one"].Name, "one"},
		{len(m2), 3},
		{m2["one"].Name, "one"},
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
		{LastN(in, 1).([]string), "[three]"},
		{LastN(in, 0).([]string), "[]"},
		{LastN(in, 2).([]string), "[two three]"},
		{LastN(in, 5).([]string), "[one two three]"},
		{LastN([]string{}, 5).([]string), "[]"},
		{LastN([]string{}, -1).([]string), "[]"},
		{LastN(records, 1).([]*record)[0].Name, "foo"},
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
		{FirstN(in, 1).([]string), "[one]"},
		{FirstN(in, 0).([]string), "[]"},
		{FirstN(in, 2).([]string), "[one two]"},
		{FirstN(in, 5).([]string), "[one two three]"},
		{FirstN([]string{}, 5).([]string), "[]"},
		{FirstN([]string{}, -1).([]string), "[]"},
	}
	for i, tc := range tests {
		has := fmt.Sprintf("%v", tc.Input)
		if tc.Want != has {
			t.Errorf("%d: want %#v (%T), was %#v (%T)", i+1, tc.Want, tc.Want, has, has)
		}
	}
}

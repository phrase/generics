package generics

import "testing"

func TestCollection(t *testing.T) {
	type record struct {
		Name  string
		Count int
	}

	list := []*record{
		{"zero", 0},
		{"one", 1},
		{"two", 2},
		{"three", 3},
	}
	c := New(list).Select(func(r *record) bool { return r.Count < 3 }).SortReverse("Count").FirstN(2)

	if v := c.Len(); v != 2 {
		t.Errorf("%d", v)
	}

	if v := c.First().(*record).Name; v != "two" {
		t.Errorf("%s", v)
	}

	if v := c.Last().(*record).Name; v != "one" {
		t.Errorf("%s", v)
	}

	m := c.FoldLeft(func(m map[string]int, v *record) map[string]int {
		m[v.Name] = v.Count * v.Count
		return m
	}).(map[string]int)
	if v := m["one"]; v != 1 {
		t.Errorf("%d", v)
	}

	if v := m["two"]; v != 4 {
		t.Errorf("%d", v)
	}
}

package generics

import (
	"reflect"
	"sort"
)

func Sort(list interface{}, fn interface{}) {
	s := sorter(list, fn)
	sort.Slice(list, s)
}

func SortReverse(list interface{}, fn interface{}) {
	s := sorter(list, fn)
	r := func(a, b int) bool {
		return !s(a, b)
	}
	sort.Slice(list, r)
}

func sorter(list interface{}, fn interface{}) func(a, b int) bool {
	v := reflect.ValueOf(list)
	//fnv := reflect.ValueOf(fn)

	etp := reflect.TypeOf(list).Elem()
	if etp.Kind() == reflect.Ptr {
		etp = etp.Elem()
	}
	_, getter := newGetter(etp, fn)
	_ = getter
	return func(a, b int) bool {
		vA := getter(v.Index(a))
		vB := getter(v.Index(b))
		switch as := vA.Interface().(type) {
		case string:
			bs := vB.Interface().(string)
			return as < bs
		case int:
			bs := vB.Interface().(int)
			return as < bs
		}
		return false
	}
}

func Map(i interface{}, fn interface{}) interface{} {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	el := t.Elem()
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	}

	tp, getter := newGetter(el, fn)
	if getter == nil {
		panic("getter is null")
	}
	out := reflect.New(reflect.SliceOf(tp))

	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		fv := getter(el)
		out.Elem().Set(reflect.Append(out.Elem(), fv))
	}
	return out.Elem().Interface()
}

func Attributes(i interface{}, name string) interface{} {
	return Map(i, name)
}

func Values(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	out := reflect.New(reflect.SliceOf(t.Elem()))
	for _, k := range v.MapKeys() {
		el := v.MapIndex(k)
		n := reflect.Append(out.Elem(), el)
		out.Elem().Set(n)
	}
	return out.Elem().Interface()
}

func Keys(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	out := reflect.New(reflect.SliceOf(t.Key()))
	for _, k := range v.MapKeys() {
		n := reflect.Append(out.Elem(), k)
		out.Elem().Set(n)
	}
	return out.Elem().Interface()
}

func Group(i interface{}, fn interface{}) interface{} {
	el := reflect.ValueOf(i).Type().Elem()

	sel := el
	if el.Kind() == reflect.Ptr {
		sel = sel.Elem()
	}

	tp, getter := newGetter(sel, fn)

	st := reflect.SliceOf(el)
	m := reflect.MakeMap(reflect.MapOf(tp, st))

	v := reflect.ValueOf(i)
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		v := getter(el)
		sl := m.MapIndex(v)
		if sl == reflect.Zero(st) || sl.Kind() == reflect.Invalid {
			sl = reflect.New(st)
			m.SetMapIndex(v, sl.Elem())
			sl = sl.Elem()
		}
		n := reflect.Append(sl, el)
		m.SetMapIndex(v, n)
	}
	return m.Interface()
}

func Index(i interface{}, fn interface{}) interface{} {
	el := reflect.ValueOf(i).Type().Elem()

	sel := el
	if el.Kind() == reflect.Ptr {
		sel = sel.Elem()
	}

	tp, getter := newGetter(sel, fn)

	m := reflect.MakeMap(reflect.MapOf(tp, el))

	v := reflect.ValueOf(i)
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		v := getter(el)
		m.SetMapIndex(v, el)
	}
	return m.Interface()
}

func Reject(i interface{}, rejecter interface{}) interface{} {
	return Select(i, negate(rejecter))
}

func Filter(i interface{}, filter interface{}) interface{} {
	return Select(i, filter)
}

func Select(i interface{}, filter interface{}) interface{} {
	fun := reflect.ValueOf(filter)
	v := reflect.ValueOf(i)
	out := reflect.New(reflect.TypeOf(i))
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		res := fun.Call([]reflect.Value{el})
		if len(res) != 1 {
			panic("must return bool")
		}
		v := res[0]
		if v.Bool() {
			n := reflect.Append(out.Elem(), el)
			out.Elem().Set(n)
		}
	}
	return out.Elem().Interface()
}

func newGetter(el reflect.Type, i interface{}) (sliceType reflect.Type, fn func(v reflect.Value) reflect.Value) {
	fnv := reflect.ValueOf(i)
	switch fnv.Kind() {
	case reflect.String:
		name := fnv.String()
		field, ok := el.FieldByName(name)
		if !ok {
			panic("no attribute with name " + name)
		}
		return field.Type, func(v reflect.Value) reflect.Value {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			return v.FieldByName(name)
		}
	case reflect.Func:
		tp := fnv.Type().Out(0)
		return tp, func(v reflect.Value) reflect.Value {
			res := fnv.Call([]reflect.Value{v})
			if len(res) != 1 {
				panic("must return one argument")
			}
			return res[0]
		}
	default:
		panic("kind " + fnv.Kind().String() + " not supported")
	}
}

func negate(fn interface{}) interface{} {
	return func(i interface{}) bool {
		v := reflect.ValueOf(fn)
		res := v.Call([]reflect.Value{reflect.ValueOf(i)})
		if len(res) != 1 {
			panic("must return bool")
		}
		return !res[0].Bool()
	}
}

func Last(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	if v.Len() == 0 {
		return nullType(v)
	}
	return v.Index(v.Len() - 1).Interface()
}

func First(i interface{}) interface{} {
	v := reflect.ValueOf(i)
	if v.Len() == 0 {
		return nullType(v)
	}
	return v.Index(0).Interface()
}

func FirstN(i interface{}, n int) interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	l := v.Len()
	if n > l {
		n = l
	}

	out := reflect.New(t)
	for idx := 0; idx < n; idx++ {
		el := v.Index(idx)
		n := reflect.Append(out.Elem(), el)
		out.Elem().Set(n)
	}

	return out.Elem().Interface()
}

func LastN(i interface{}, n int) interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	l := v.Len()
	if n > l {
		n = l
	}
	out := reflect.New(t)
	for i := 0; i < n; i++ {
		idx := i + (l - n)
		el := v.Index(idx)
		n := reflect.Append(out.Elem(), el)
		out.Elem().Set(n)
	}

	return out.Elem().Interface()
}

func nullType(v reflect.Value) interface{} {
	et := v.Type().Elem()
	if et.Kind() == reflect.Ptr {
		return nil
	}
	return reflect.New(et).Elem().Interface()
}

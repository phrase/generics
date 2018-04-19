package generics

import (
	"reflect"
)

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

func Index(i interface{}, fn interface{}) interface{} {
	el := reflect.ValueOf(i).Type().Elem()
	m := reflect.MakeMap(reflect.MapOf(reflect.TypeOf("test"), el))
	fun := reflect.ValueOf(fn)
	v := reflect.ValueOf(i)
	for i := 0; i < v.Len(); i++ {
		el := v.Index(i)
		res := fun.Call([]reflect.Value{el})
		if len(res) != 1 {
			panic("expected one string to return")
		}
		m.SetMapIndex(res[0], el)
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

func Head(i interface{}, n int) interface{} {
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

func Tail(i interface{}, n int) interface{} {
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

package generics

import "reflect"

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

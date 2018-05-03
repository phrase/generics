package generics

import (
	"reflect"
	"strings"
)

func (c *Collection) Join(mapOrList interface{}) *Collection {
	t := reflect.TypeOf(mapOrList)
	name := nameFromMap(t)
	return c.JoinGeneric(mapOrList, name, name+"ID", "ID")
}

func (c *Collection) JoinGeneric(mapOrList interface{}, name, foreignKeyName, primaryKeyName string) *Collection {
	t := reflect.TypeOf(mapOrList)

	mv := reflect.ValueOf(mapOrList)
	sv := reflect.ValueOf(c.collection)
	switch t.Kind() {
	case reflect.Slice:
		st := t.Elem()
		if st.Kind() == reflect.Ptr {
			st = st.Elem()
		}
		f, _ := st.FieldByName(primaryKeyName)
		newMV := reflect.MakeMap(reflect.MapOf(f.Type, t.Elem()))

		for i := 0; i < mv.Len(); i++ {
			el := mv.Index(i)
			var key reflect.Value
			if el.Kind() == reflect.Ptr {
				key = el.Elem().FieldByName(primaryKeyName)
			} else {
				key = el.FieldByName(primaryKeyName)
			}
			newMV.SetMapIndex(key, el)
		}
		mv = newMV
	case reflect.Map:
		// ok
	default:
		panic("expected second argument to be map, was " + t.Kind().String())
	}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		account := mv.MapIndex(v.FieldByName(foreignKeyName))
		v.FieldByName(name).Set(account)
	}
	return c
}

func nameFromMap(in reflect.Type) string {
	el := strings.Split(in.Elem().String(), ".")
	return el[len(el)-1]
}

type Payment struct {
	ID        int      `json:"id"`
	AccountID int      `json:"account_id"`
	Account   *Account `json:"account"`
}

type Account struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

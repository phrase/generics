package generics

import "reflect"

func New(i interface{}) *Collection {
	return &Collection{i}
}

type Collection struct {
	collection interface{}
}

func (c *Collection) FoldLeft(folder interface{}) interface{} {
	return FoldLeft(c.collection, folder)
}

func (c *Collection) Values(folder interface{}) interface{} {
	return Values(c.collection)
}

func (c *Collection) Map(mapper interface{}) *Collection {
	return New(Map(c.collection, mapper))
}

func (c *Collection) Select(fn interface{}) *Collection {
	return New(Select(c.collection, fn))
}

func (c *Collection) Reject(fn interface{}) *Collection {
	return New(Reject(c.collection, fn))
}

func (c *Collection) Group(fn interface{}) interface{} {
	return Group(c.collection, fn)
}

func (c *Collection) Index(fn interface{}) interface{} {
	return Index(c.collection, fn)
}

func (c *Collection) First() interface{} {
	return First(c.collection)
}

func (c *Collection) Last() interface{} {
	return Last(c.collection)
}

func (c *Collection) FirstN(n int) *Collection {
	return New(FirstN(c.collection, n))
}

func (c *Collection) LastN(n int) *Collection {
	return New(LastN(c.collection, n))
}

func (c *Collection) Cast() interface{} {
	return c.collection
}

func (c *Collection) Sort(fn interface{}) *Collection {
	Sort(c.collection, fn)
	return c
}

func (c *Collection) SortReverse(fn interface{}) *Collection {
	SortReverse(c.collection, fn)
	return c
}

func (c *Collection) Len() int {
	return reflect.ValueOf(c.collection).Len()
}

package main

type element struct {
	value int
}

func newElement(value int) *element {
	e := new(element)
	e.value = value
	return e
}

func (e element) getValue() int {
	return e.value
}

func (e *element) setValue(newValue int) {
	e.value = newValue
}

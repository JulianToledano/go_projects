package main

type element struct {
	content int
}

func (e element) getContent() int {
	return e.content
}

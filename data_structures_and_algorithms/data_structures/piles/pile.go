package main

type pile struct {
	root *node
}

func (p *pile) push(e element) {
	n := new(node)
	n.setElement(e)
	p.insert(n)
}
func (p *pile) insert(n *node) {
	n.next = p.root
	p.root = n
}

func (p *pile) pop() element {
	firstElement := p.root
	p.root = p.root.getNext()
	return firstElement.getElement()
}

package inutil

type Thing struct {
	Name string
}

func (n *Thing) getName() string {
	return n.Name
}

func (n *Thing) copy() *Thing {
	r := &Thing{n.Name}
	return r
}

// Equality compare: same as value == value
func (this *Thing) EqualsTo(other *Thing) bool {
	return this.Name == other.Name
	// Output: true is they have the same values
}

// Same reference
// Equality compare: same as *ref ==  *ref
func (this *Thing) Equals(other *Thing) bool {
	return this == other
	// Output: true is they are the same "instance?"
}
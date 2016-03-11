package inutil

type Thing struct {
	Name string
}

func (thing *Thing) IsEmpty() bool {
	return thing.Name == ""
}

func (thing *Thing) IsNilOrEmpty() bool {
	return thing == nil || thing.Name == ""
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

func IsThingNamed(name string) func(thing *Thing) bool {
	f := func(thing *Thing) bool {
		return thing.Name == name
	}
	return f
}

type Things []*Thing;

func MakeThings(thing ...*Thing) Things {
	return thing
}

func (this Things) FirstOrDefault(find func(t *Thing) bool) *Thing {
	for _, thing := range this[:] {
		if find(thing) {
			return thing
		}
	}
	return nil
}
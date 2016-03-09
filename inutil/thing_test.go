package inutil

import "testing"

// Get Equality Right
func Test_equality(t *testing.T) {

	thing := &Thing{"A"}
	other := &Thing{"A"}
	odd := &Thing{"B"}

	if thing == other {
		t.Error("Wtf?")
	}
	if thing == odd {
		t.Error("Wtf?")
	}
	if !thing.EqualsTo(other) {
		t.Error("WTF?")
	}
	if thing.Equals(other) {
		t.Error("WTF?")
	}
	if thing.Equals(odd) {
		t.Error("WTF?")
	}
	if thing.EqualsTo(odd) {
		t.Error("Wtf")
	}
	if thing != thing {
		t.Error("Wtf?")
	}

	if !thing.Equals(thing) {
		t.Error("wtf?")
	}
	// value compare : b'cause right its an interface
	if Anything(thing) != thing {
		t.Error("Wtf?")
	}
}

func Test_GotTypeAssertions(t *testing.T) {
	var anything Anything = Thing{"A"}  // anything has dynamic Thing int and value {"A"}
	thing := anything.(Thing)           // thing has type Thing and value {"A"}
	t.Logf("thing: %v == anything: %v => %v ", thing, anything, thing == anything)
	var thingBack = anything.(Thing)
	t.Logf("thing: %v Equals anythng: %v => %v ", thing, anything, thing.Equals(&thingBack))
	t.Logf("thing: %v EqualsTo anything: %v => %v ", thing, anything, thing.EqualsTo(&thingBack))
}

type Anything interface {

}
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

//
// Something
//
func Test_GotTypeAssertions(t *testing.T) {
	/*
		Something Else
	*/
	var anything Anything = Thing{"A"}  // anything has dynamic Thing int and value {"A"}
	thing := anything.(Thing)           // thing has type Thing and value {"A"}
	t.Logf("thing: %v == anything: %v => %v ", thing, anything, thing == anything)
	var thingBack = anything.(Thing)
	t.Logf("thing: %v Equals anythng: %v => %v ", thing, anything, thing.Equals(&thingBack))
	t.Logf("thing: %v EqualsTo anything: %v => %v \n\n", thing, anything, thing.EqualsTo(&thingBack))
}

func Test_Something(t *testing.T) {
	var anything Anything = Thing{"A"}
	if anything.(Thing).Name != "A" {
		t.Error("Wtf?")
	}
}

func Test_Again(t *testing.T) {
	var anything Anything = Thing{"A"}
	if thing, ok := anything.(Thing); ok && thing.Name != "A" {
		t.Error("Wtf?")
	} else if ok && thing.Name == "" {
		t.Error("Wtf?")
	} else if ok {
		t.Log("Ok \n")
	} else {
		t.Errorf("Not Ok %v", anything)
	}
}

func Test_More(t *testing.T) {

	var anything Anything = "x"

	switch v := anything.(type) {

	case string:
		t.Logf("v<string>: %v", v)
	// look Ma' No breaks
	case Anything:
		t.Error("SHould've break!")
	case int32, int64:
		t.Error("WTF?")
	case Thing:
		t.Error("WTF")
	default:
		t.Error("what?")
	}

}

func Test_Things(t *testing.T) {
	var things [1] Anything
	if thing := things[0]; thing != nil {
		t.Errorf("What a f? %v", thing)
	}
	things[0] = Thing{"X"}
	if things[0].(Thing).Name != "X" {
		t.Error("Really?")
	}
}

func Test_MoreThings(t *testing.T) {
	things := [3]Anything{Thing{"X"}, Thing{"Y"}, "no"}
	for i := 0; i < len(things); i++ {
		thing := things[i]
		switch thing.(type) {
		case Thing:
			thing := thing.(Thing)
			if thing.Name == "" {
				t.Error("No such thing \n\n")
			}else {
				t.Logf("d'thing %v", thing.Name)
			}
		case Anything:
			t.Logf("%v", thing)
		}
	}
}

func IsThingNamedX(anything Anything) bool {
	thing, ok := anything.(Thing)
	return (ok && thing.Name == "X")
}

func Test_FirstOrDefault(t *testing.T) {
	things := []Anything{Thing{"X"}}
	thing := FirstOrDefault(things, IsThingNamedX)
	if thing == nil {
		t.Error("Not Found")
	}else {
		t.Logf("found: %s", thing.(Thing).Name)
	}
	things = []Anything{Thing{"Y"}}
	thing = FirstOrDefault(things, IsThingNamedX)
	if thing != nil {
		t.Error("???")
	}
	var _things Things = MakeThings(&Thing{"Z"})
	if _things.FirstOrDefault(IsThingNamed("Z")).IsEmpty() {
		t.Error("Wrong")
	}
	if !_things.FirstOrDefault(IsThingNamed("X")).IsNilOrEmpty() {
		t.Error("Wrong")
	}

}

func Test_IncrementNumbers(t *testing.T) {

	numbersSource := []int{0, 1, 2}

	increment := func(x *int) {
		*x++
	}

	// Increment number in numbers ,
	// while numbers is a value, a copy
	// of the original slice,
	// slice elements are pointers to the original value
	// so.. original slice/array's elements will
	// change... mutate accordingly
	increments := func(numbers []int) {
		for i := 0; i < len(numbers); i++ {
			increment(&numbers[i])
		}
	}

	increments(numbersSource)

	if numbersSource[0] != 1 {
		t.Error("DOesn't work")
	}else {
		t.Log(numbersSource)
	}
}








package inutil

import (
	"testing"
	"time"
)

const ok = ":)"
const notOk = ":("

func Test_Routinas(t *testing.T) {

	var result = ""

	messages := make(chan string)

	// anonymous func
	go func() {
		//
		messages <- ReverseString("pin")
	}()
	// await<string> ;)
	result = <-messages

	if result != "nip" {
		t.Error("Doesn't wait")
	}

	// ------------------------------------------------------------------------
	// NO WAIT:
	result = ""
	// NoChannel, fire and forget
	go func() {
		//  No Wait
		result = ReverseString("on")
	}()
	// NO WAIT
	//result = <- messages
	if result != "" {
		t.Errorf("Doesn't work as I think it should, result=&v", result)
	}

	wait := make(chan Anything)
	result = notOk
	go func() {
		result = ok
		wait <- Thing{}
	}()
	<-wait
	if isNotOk(result) {
		t.Error(result)
	}

	// NOWAIT
	result = notOk
	go func() {
		result = ok
		wait <- Thing{}
	}()
	// <- wait
	if isOk(result) {
		t.Error(result)
	}

	// wait ...
	getOk := func() string {
		time.Sleep(time.Second * 2)
		return ":)"
	}
	receiveOk := func(ok string) {
		result = ok
	}
	result = notOk
	WaitFor(getOk, receiveOk)
	if isNotOk(result) {
		t.Error(result)
	}

	// more waiting ..
	done := make(chan bool)
	waitfor := func(done chan bool, callback func()) {
		time.Sleep(time.Second)
		done <- true
		callback()
	}
	result = notOk
	go waitfor(done, func() {
		result = ok
	})
	<-done
	if isNotOk(result) {
		t.Error(notOk)
	}
	// and wait again ...
	actioned := false
	calledBack := false
	WaitForAction(
		// Anonymous Action
		func() {
			time.Sleep(time.Second)
			actioned = true;
			//			t.Logf("action...")
		},
		// anonymous Callback
		func() {
			calledBack = true;
			//			t.Logf("Callback")
		},
	)

	if !actioned || !calledBack {
		t.Errorf("Actioned: %v, Calledback: %v   %v", actioned, calledBack, notOk)
	}

	// ...Then
	var sleep Action = func() {
		time.Sleep(time.Second)
	}

	isDone := false

	setIsDone := func() {
		isDone = true
	}

	sleep.Then(setIsDone)
	if !isDone {
		t.Error("is Not DOne")
	}

}

func isOk(s string) bool {
	return s == ok
}
func isNotOk(s string) bool {
	return s == notOk
}


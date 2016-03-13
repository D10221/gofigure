package main

import (
	"fmt"
	"time"
	"github.com/D10221/gofigure/inutil"
)

func main() {
	done := make(chan bool)
	actioned := false
	calledBack := false
	inutil.WaitForAction(done,
		// Anonymous Action
		func() {
			time.Sleep(time.Second)
			actioned = true;
			fmt.Print("action...")
		},
		// anonymous Callback
		func() {
			calledBack = true;
			fmt.Print("Callback")
		},
	)
	<-done
	if !actioned || !calledBack {
		fmt.Print("Actioned: %v, Calledback: %v   %v", actioned, calledBack, ":)")
	}

}
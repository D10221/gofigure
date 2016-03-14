package inutil

func WaitThen(action Action, callback Action) {
	done := make(chan bool)
	go func() {
		action()
		done <- true
	}()
	<-done
	close(done)
	callback()
}

func WaitFor(action Call, callback Callback) {
	channel := make(chan string)
	go func() {
		action()
		channel <- action()
		close(channel)
	}()
	callback(<-channel)
}

type Action func();

type Call func() string;

type Callback func(s string);

func (action Action) Then(then Action) {

	done := make(chan bool)

	defer func() {
		<-done
		close(done)
	}()

	go func() {
		action()
		done <- true
	}()

	then()
}

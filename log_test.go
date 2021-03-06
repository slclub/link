package link

import (
	"testing"
)

func TestLog(t *testing.T) {
	f1 := func() {
		INFO("I am an info level log.")
		DEBUG("I am an debug level log.")
		WARN("I am an warn level log.")
		ERROR("this an error log. so you should be careful.")
	}

	for i := 0; i < 100; i++ {
		f1()
	}

	var ch_out = make(chan int, 1)
	ch_out <- 1

	for {
		select {
		case <-ch_out:
			t.Log("OUT WITH RETURN")
			return
		}
	}

}

func TestDebug_print(t *testing.T) {
	DEBUG_PRINT("debug", " nothing")
}

package sewer

import "testing"

func TestLogger(t *testing.T) {
	var gotMsg *Message
	var wait = make(chan bool)
	AddFilter("hello.*", func(msg *Message) {
		gotMsg = msg
		wait <- true
	})
	AddFilter("should.never.happen", func(msg *Message) {
		t.Errorf("This filter should never be called")
	})
	Log("hello.world", "foo", "bar")
	<-wait
	if gotMsg == nil {
		t.Errorf("Expected to receive a message on hello.* filter")
		return
	}
	if gotMsg.Event != "hello.world" {
		t.Errorf("Expected to have hello event, got: %s", gotMsg.Event)
	}
	if foo, ok := gotMsg.Payload["foo"]; !ok || foo.(string) != "bar" {
		t.Errorf("Expected to have foo = bar, got foo = %v", foo)
	}
}

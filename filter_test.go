package sewer

import (
	"testing"
)

func TestFilter_Match(t *testing.T) {
	f := &attachedFilter{nil, "hello.*"}
	if !f.Match("hello.world") {
		t.Errorf("Expected to match hello.world filter")
	}
	if f.Match("bye.bye") {
		t.Errorf("Expected to not match bye.bye filter")
	}
}
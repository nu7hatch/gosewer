package iolog

import (
	sewer "github.com/nu7hatch/gosewer"
	"time"
)

func ExampleFilter() {
	sewer.AddFilter("hello", Filter)
	sewer.Log("hello", "foo", "bar")
}
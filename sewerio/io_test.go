package sewerio

import (
	sewer "github.com/nu7hatch/gosewer"
	"os"
)

func ExampleFilter() {
	sewer.AddFilter("hello", NewFilter(os.Stderr))
	sewer.Log("hello", "foo", "bar")
}

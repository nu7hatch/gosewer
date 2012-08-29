package fmtlog

import (
	sewer "github.com/nu7hatch/gosewer"
	"time"
)

func ExampleFmtLog() {
	sewer.AddFilter("hello", FmtLog)
	sewer.Log("hello", "foo", "bar")
}
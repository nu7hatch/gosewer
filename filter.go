package sewer

import (
	"path"
)

type Filter func(*Message)

type attachedFilter struct {
	Filter
	Pattern string
}

func (self *attachedFilter) Match(event string) (ok bool) {
	var err error
	if ok, err = path.Match(self.Pattern, event); err != nil {
		panic(err.Error())
	}
	return
}

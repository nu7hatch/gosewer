package sewer

type Logger struct {
	ch      chan *Message
	quit    chan bool
	filters []*attachedFilter
}

func NewLogger() (log *Logger) {
	return NewBufferedLogger(0)
}

func NewBufferedLogger(bufferSize int) (log *Logger) {
	log = &Logger{
		make(chan *Message, bufferSize),
		make(chan bool),
		[]*attachedFilter{},
	}
	go log.run()
	return
}

func (self *Logger) run() {
	for {
		select {
		case <-self.quit:
			return
		case msg := <-self.ch:
			self.deliver(msg)
		}
	}
}

func (self *Logger) deliver(msg *Message) {
	for _, filter := range self.filters {
		if filter.Match(msg.Event) {
			filter.Filter(msg)
		}
	}
}

func (self *Logger) Quit() {
	self.quit <- true
}

func (self *Logger) AddFilter(pattern string, filter Filter) {
	self.filters = append(self.filters, &attachedFilter{filter, pattern})
}

func (self *Logger) Log(event string, argv ...interface{}) {
	self.ch <- NewMessage(event, argv[0].([]interface{}))
}
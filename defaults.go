package sewer

var DefaultLogger *Logger = NewLogger()

func Log(event string, args ...interface{}) {
	DefaultLogger.Log(event, args)
}

func AddFilter(pattern string, filter Filter) {
	DefaultLogger.AddFilter(pattern, filter)
}
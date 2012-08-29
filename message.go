package sewer

import "time"

type Message struct {
	Event   string
	Payload map[string]interface{}
	Stamp   time.Time
}

func NewMessage(event string, args []interface{}) *Message {
	argc := len(args)
	i, payload := 0, make(map[string]interface{})
	if argc % 2 == 0 {
		for i < argc {
			payload[args[i].(string)] = args[i+1]
			i += 2
		}
	}
	return &Message{event, payload, time.Now()}
}
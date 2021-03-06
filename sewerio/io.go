package sewerio

import (
	"encoding/json"
	"fmt"
	sewer "github.com/nu7hatch/gosewer"
	"io"
	"time"
)

func NewFilter(writer io.Writer) sewer.Filter {
	return func(msg *sewer.Message) {
		payload, _ := json.Marshal(msg.Payload)
		stamp := msg.Stamp.Format(time.Stamp)
		s := fmt.Sprintf("%s %s -> %s\n", stamp, msg.Event, payload)
		writer.Write([]byte(s))
	}
}

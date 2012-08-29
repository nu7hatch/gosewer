package fmtlog

import (
	"encoding/json"
	"fmt"
	sewer "github.com/nu7hatch/gosewer"
	"time"
)

func FmtLog(msg *sewer.Message) {
	payload, _ := json.Marshal(msg.Payload)
	stamp := msg.Stamp.Format(time.Stamp)
	fmt.Printf("%s %s -> %s\n", stamp, msg.Event, payload)
}

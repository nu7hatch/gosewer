package sewerredis

import (
	sewer "github.com/nu7hatch/gosewer"
	redis "github.com/simonz05/godis"
	"testing"
	"time"
	"encoding/json"
)

var client *redis.Client

func init() {
	client = redis.New("tcp:127.0.0.1:6379", 0, "")
	client.Flushdb()
}

func TestFilter(t *testing.T) {
	sewer.AddFilter("hello", NewFilter(client))
	sewer.Log("hello", "foo", "bar")
	<-time.After(100 * time.Millisecond)
	rep, err := client.Zrange("hello", 0, -1)
	if err != nil {
		t.Errorf("Expected to get list of stored logs, got error: %v", err)
		return
	}
	lines := rep.BytesArray()
	if len(lines) != 1 {
		t.Errorf("Expected to have one log item, got %s", len(lines))
		return
	}
	var data map[string]interface{}
	if err := json.Unmarshal(lines[0], &data); err != nil {
		t.Errorf("Expected to unmarshal payload properly, got error: %v", err)
		return
	}
	if payload, ok := data["payload"].(map[string]interface{}); !ok || payload["foo"] != "bar" {
		t.Errorf("Expected to have proper payload, got %v", payload)
	}
}

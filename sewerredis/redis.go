package sewerredis

import (
	"encoding/json"
	"errors"
	sewer "github.com/nu7hatch/gosewer"
	redis "github.com/simonz05/godis"
	"net/url"
	"regexp"
	"strconv"
)

func NewFilter(clientOrUri interface{}) sewer.Filter {
	if client, ok := clientOrUri.(*redis.Client); ok {
		return newFilterForClient(client)
	} else if uri, ok := clientOrUri.(string); ok {
		return newFilterForUri(uri)
	}

	panic("Param must be either a Redis client or database URL")
	return nil
}

func newFilterForClient(client *redis.Client) sewer.Filter {
	return func(msg *sewer.Message) {
		if payload, err := json.Marshal(msg); err == nil {
			client.Zadd(msg.Event, float64(msg.Stamp.Unix()), payload)
		}
	}
}

func getRedisDb(path string) (res int, err error) {
	var dbFinder = regexp.MustCompile("\\/(\\d+)")
	res = -1

	if found := dbFinder.FindString(path); found != "" {
		res, _ = strconv.Atoi(found)
	} else {
		err = errors.New("database not specified")
	}

	return
}

func parseRedisUrl(uri string) (res *url.URL, err error) {
	if res, err = url.Parse(uri); err == nil {
		if res.Scheme != "redis" {
			err = errors.New("invalid database address")
		}
	}
	return
}

func newFilterForUri(uri string) sewer.Filter {
	var err error
	var url *url.URL
	var password string
	var db int
	var client *redis.Client

	if url, err = parseRedisUrl(uri); err != nil {
		panic("Invalid redis URL: " + uri)
	}
	if db, err = getRedisDb(url.Path); err != nil {
		panic("Invalid database: " + url.Path)
	}
	if url.User != nil {
		password, _ = url.User.Password()
	}

	client = redis.New("tcp:"+url.Host, db, password)
	return newFilterForClient(client)
}

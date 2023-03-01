package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"strings"
	"time"
)

type Option func(*options)

type options struct {
	servers        []string
	sessionTimeout time.Duration
	otherOpts      []func(conn *zk.Conn)
}

func setDefault() options {
	return options{
		servers:        []string{},
		sessionTimeout: time.Second * 6,
	}
}

func SetServers(servers string) Option {
	return func(o *options) {
		if strings.IndexAny(servers, ";") >= 0 {
			o.servers = strings.Split(servers, ";")
		} else if strings.IndexAny(servers, ",") >= 0 {
			o.servers = strings.Split(servers, ",")
		} else {
			o.servers = []string{servers}
		}
	}
}

func SetSessionTimeout(sessionTimeout int64) Option {
	return func(o *options) {
		if sessionTimeout <= 0 {
			return
		}
		_timeout := int64(time.Millisecond) * sessionTimeout
		o.sessionTimeout = time.Duration(_timeout)
	}
}

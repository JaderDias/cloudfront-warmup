package warmup

import (
	"net"
)

type NetLookup interface {
	LookupIP(host string) ([]net.IP, error)
}

type MyNetLookup struct {
}

func (l *MyNetLookup) LookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}

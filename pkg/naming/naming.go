package naming

import (
	"context"
)

// metadata common key
const (
	MetaColor   = "color"
	MetaWeight  = "weight"
	MetaCluster = "cluster"
	MetaZone    = "zone"
)

// Instance represents a server the client connects to.
type Instance struct {
	// Region bj/sh/gz
	Region string `json:"region"`
	// Zone is IDC.
	Zone string `json:"zone"`
	// Env prod/pre、uat/fat1
	Env string `json:"env"`
	// AppID is mapping servicetree appid.
	AppID string `json:"appid"`
	// Hostname is hostname from docker.
	Hostname string `json:"hostname"`
	// Addrs is the adress of app instance
	// format: scheme://host
	Addrs []string `json:"addrs"`
	// Version is publishing version.
	Version string `json:"version"`
	// LastTs is instance latest updated timestamp
	LastTs int64 `json:"latest_timestamp"`
	// Metadata is the information associated with Addr, which may be used
	// to make load balancing decision.
	Metadata map[string]string `json:"metadata"`
	Status   int64
}

// App is the data
type App struct {
	Instances map[string][]*Instance `json:"zone_instances"`
	LastTs    int64                  `json:"latest_timestamp"`
	Scheduler *Scheduler             `json:"scheduler"`
	Err       string                 `json:"err"`
}

type Scheduler struct {
	Clients map[string]*ZoneStrategy `json:"clients"`
}

// ZoneStrategy is the scheduling strategy of all zones
type ZoneStrategy struct {
	Zones map[string]*Strategy `json:"zones"`
}

// Strategy is zone scheduling strategy.
type Strategy struct {
	Weight int64 `json:"weight"`
}

// Resolver resolve naming service
type Resolver interface {
	Fetch(context.Context) (map[string][]*Instance, bool)
	//Unwatch(id string)
	Watch() <-chan struct{}
	Close() error
}

// Registry Register an instance and renew automatically
type Registry interface {
	Register(context.Context, *Instance) (context.CancelFunc, error)
	Close() error
}

// Builder resolver builder.
type Builder interface {
	Build(id string, options ...BuildOpt) Resolver
	Scheme() string
}

package driver

import "time"

// Client has Data store's client interface.
// This is the interface of the data source client used by Content-Management-API.
// If you want to use a data source client that is not provided, you need to prepare a client implementation that implements the interface.
type Client interface {
	Connect() (Client, error)
	Disconnect() error

	// Monitor embedded interface
	Monitor
}

// Monitor is an interface for checking communication with the data store at regular intervals.
type Monitor interface {
	Ping(duration time.Duration) error
	StartWatch()
	StopWatch()
	Watch(stopCh chan struct{}) error
}

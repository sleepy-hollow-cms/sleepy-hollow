package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	uri               string
	maxConnectionPool uint64
	timeOut           time.Duration

	cache    sync.Map
	cacheKey string

	// for monitor
	watchDuration time.Duration
	stopCh        chan struct{}
}

func NewClient() driver.Client {
	return &Client{
		uri:               fmt.Sprintf("mongodb://%s:%s@%s:%v", config.Conf.MongoDB.User, config.Conf.MongoDB.Password, config.Conf.MongoDB.Host, config.Conf.MongoDB.Port),
		maxConnectionPool: 100,
		timeOut:           10 * time.Second,
		cacheKey:          "mongo-client",
		watchDuration:     5 * time.Second,
	}
}

func (c *Client) Connect() (driver.Client, error) {
	option := options.Client().ApplyURI(c.uri)
	option.SetMaxPoolSize(c.maxConnectionPool)

	client, err := mongo.NewClient(option)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeOut)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	c.cache.Store(c.cacheKey, client)

	return c, nil
}

func (c *Client) Disconnect() error {
	client, err := c.load()

	if err != nil {
		return err
	}

	err = client.Disconnect(context.Background())

	if err != nil {
		return err
	}

	c.cache.Delete(c.cacheKey)
	return nil
}

func (c *Client) load() (*mongo.Client, error) {
	load, ok := c.cache.Load(c.cacheKey)
	if !ok {
		return nil, errors.New("could not load client by cache")
	}
	return load.(*mongo.Client), nil
}

func (c *Client) Get() (*mongo.Client, error) {
	client, err := c.load()
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}
	return client, nil
}

func (c *Client) StartWatch() {
	stopCh := make(chan struct{})
	c.stopCh = stopCh
	go c.Watch(stopCh)
}

func (c *Client) StopWatch() {
	close(c.stopCh)
}

func (c *Client) Watch(stopCh chan struct{}) error {
	ticker := time.NewTicker(c.watchDuration)
	defer ticker.Stop()
	for {
		select {
		case <-stopCh:
			log.Println("stop watching mongo-client.")
			return nil
		case <-ticker.C:
			err := c.Ping(2 * time.Second)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Client) Ping(duration time.Duration) error {
	db, err := c.load()
	if err != nil {
		return errors.New("cant db db")
	}

	timeout, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	err = db.Ping(timeout, readpref.Primary())

	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}

	return nil
}

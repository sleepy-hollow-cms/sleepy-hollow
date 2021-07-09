package mongo

import (
	"content-management-api/env"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Client struct {
	uri               string
	maxConnectionPool uint64
	timeOut           time.Duration

	// cache mechanism. I made it loosely.
	mu          sync.Mutex
	client      *mongo.Client
	clientCache map[string]*mongo.Client
	cacheKey    string
}

func NewClient() *Client {
	config := env.GetMongoConfig()

	return &Client{
		uri:               fmt.Sprintf("mongodb://%s:%s@%s:%v", config.User, config.Password, config.Host, config.Port),
		maxConnectionPool: 100,
		timeOut:           30 * time.Second,
		clientCache:       make(map[string]*mongo.Client),
		cacheKey:          "mongo-client",
	}
}

func (d *Client) Connect() (*Client, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	option := options.Client().ApplyURI(d.uri)

	option.SetMaxPoolSize(d.maxConnectionPool)

	client, err := mongo.NewClient(option)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), d.timeOut)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	d.clientCache[d.cacheKey] = client

	return d, nil
}

func (d *Client) Close() *Client {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.client != nil {
		d.client.Disconnect(context.Background())
		delete(d.clientCache, d.cacheKey)
		d.client = nil
	}
	return d
}

func (d *Client) Get() *mongo.Client {
	return d.clientCache[d.cacheKey]
}

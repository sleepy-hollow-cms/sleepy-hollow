package mongo

import (
	"content-management-api/env"
	"context"
	"errors"
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

	cache    sync.Map
	cacheKey string
}

func NewClient() *Client {
	config := env.GetMongoConfig()
	return &Client{
		uri:               fmt.Sprintf("mongodb://%s:%s@%s:%v", config.User, config.Password, config.Host, config.Port),
		maxConnectionPool: 100,
		timeOut:           30 * time.Second,
		cacheKey:          "mongo-client",
	}
}

func (d *Client) Connect() (*Client, error) {
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

	d.cache.Store(d.cacheKey, client)

	return d, nil
}

func (d *Client) Close() error {
	client, err := d.load()

	if err != nil {
		return err
	}

	client.Disconnect(context.Background())
	d.cache.Delete(d.cacheKey)
	return nil
}

func (d *Client) load() (*mongo.Client, error) {
	load, ok := d.cache.Load(d.cacheKey)
	if !ok {
		return nil, errors.New("could not load client by cache")
	}
	return load.(*mongo.Client), nil
}

func (d *Client) Get() (*mongo.Client, error) {
	client, err := d.load()
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}
	return client, nil
}

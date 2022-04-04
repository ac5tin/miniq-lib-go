package main

import (
	"context"
	"miniq/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
}

// creates a new miniq connection client
func NewClient(address string) (*Client, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(address, opts, grpc.WithBlock())
	if err != nil {
		return nil, nil
	}

	client := new(Client)
	client.conn = cc
	return client, nil
}

// closes the connection to the miniq server
func (c *Client) Close() error {
	return c.conn.Close()
}

// Add new task
func (c *Client) AddTask(channel *string, data *[]byte) error {
	client := proto.NewMiniQClient(c.conn)
	// prepare request
	req := new(proto.AddTaskRequest)
	req.Channel = *channel
	req.Data = *data
	// send request
	if _, err := client.AddTask(context.Background(), req); err != nil {
		return err
	}
	return nil
}

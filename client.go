package miniq

import (
	"context"
	"io"
	"log"

	"github.com/ac5tin/miniq-lib-go/proto"

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
func (c *Client) AddTask(channel string, data *[]byte) error {
	client := proto.NewMiniQClient(c.conn)
	// prepare request
	req := new(proto.AddTaskRequest)
	req.Channel = channel
	req.Data = *data
	// send request
	if _, err := client.AddTask(context.Background(), req); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTasks(channel string, status TaskStatus) (<-chan *Task, error) {
	client := proto.NewMiniQClient(c.conn)
	// prepare request
	req := new(proto.GetTaskRequest)
	req.Status = proto.TaskStatus(status)
	req.Channel = channel
	stream, err := client.GetTasks(context.Background(), req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *Task)
	ctx := stream.Context()
	go func() {
		for {
			t, err := stream.Recv()
			// error handling
			{
				if err == io.EOF {
					// EOF = end of stream
					close(ch)
					return
				}
				if err != nil {
					log.Printf("[client] Failed to get tasks, Err: %v\n", err) // debug
					close(ch)
					return
				}
			}
			// handle nil
			{
				if t == nil {
					continue
				}
			}
			// prepare task data
			ch <- ProtoTask2Task(t)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Printf("[client] %v\n", err.Error()) // debug
		}
		// close if not already closed
		if _, ok := <-ch; ok {
			close(ch)
		}
	}()

	return ch, err
}

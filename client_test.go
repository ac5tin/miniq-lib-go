package miniq

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c, err := NewClient("localhost:8080")
	if err != nil {
		t.Error(err)
	}

	t.Log("Successfully connected to server")

	// add task
	data := make([]byte, 100000000) // create big data
	t.Run("AddTask", func(t *testing.T) {
		// data := []byte("test")
		rand.Read(data)
		if err := c.AddTask("test", &data); err != nil {
			t.Error(err)
		}
	})

	//get task
	{
		ch, err := c.GetTasks("test", TaskStatusPending)
		if err != nil {
			t.Error(err)
		}
		task := <-ch
		t.Logf("Received task with ID: %s and data size:  %d\n", task.ID, len(task.Data))
		assert.Equal(t, len(data), len(task.Data))
		assert.Equal(t, data, task.Data)
	}
}

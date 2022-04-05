package miniq

import "testing"

func TestClient(t *testing.T) {
	c, err := NewClient("localhost:8080")
	if err != nil {
		t.Error(err)
	}

	t.Log("Successfully connected to server")

	// add task
	t.Run("AddTask", func(t *testing.T) {
		data := []byte("test")
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
		t.Logf("Received task with ID: %s and data: %s\n", task.ID, task.Data)
	}
}

package main

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
}

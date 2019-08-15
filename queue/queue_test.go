package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	var newQueue Topic
	actualQueue := newQueue.New("TestTopic")
	expectedQueue := Topic{"TestTopic", []Item{}}

	if actualQueue.TopicName != expectedQueue.TopicName {
		t.Errorf("Expected topic name: 'TestTopic' got '%s'", actualQueue.TopicName)
	}

	if len(actualQueue.Queue) != len(expectedQueue.Queue) {
		t.Errorf("Expected queue length: %d got %d", len(actualQueue.Queue), len(expectedQueue.Queue))
	}

}

func TestGet(t *testing.T) {
	var newTopic Topic
	actualTopic := newTopic.New("TestTopic")
	actualTopic.Insert("This is a test message")
	actualItem := actualTopic.Get()

	expectedTopic := Topic{"TestTopic", []Item{}}
	expectedTopic.Queue = append(expectedTopic.Queue, Item{1, "This is a test message", "", "TestTopic"})
	expectedItem := expectedTopic.Queue[0]

	if (actualItem.ID != expectedItem.ID) || (actualItem.Message != expectedItem.Message) || (actualItem.Topic != expectedItem.Topic) {
		t.Errorf("Item struct payload is invalid.")
	}
}

func TestDelete(t *testing.T) {
	var newTopic Topic
	actualTopic := newTopic.New("TestTopic")
	actualTopic.Insert("This is a test message for deletion")
	actualTopic.Delete()

	expectedTopic := Topic{"TestTopic", []Item{}}

	if len(actualTopic.Queue) != len(expectedTopic.Queue) {
		t.Errorf("Expected queue length: %d got %d", len(actualTopic.Queue), len(expectedTopic.Queue))
	}

}

func TestInsert(t *testing.T) {
	var newTopic Topic
	actualTopic := newTopic.New("TestTopic")
	actualTopic.Insert("This is a test message")
	actualQueue := actualTopic.Queue

	expectedTopic := Topic{"TestTopic", []Item{}}
	expectedTopic.Queue = append(expectedTopic.Queue, Item{1, "This is a test message", "", "TestTopic"})
	expectedQueue := expectedTopic.Queue

	if len(actualTopic.Queue) != len(expectedTopic.Queue) {
		t.Errorf("Expected queue length: %d got %d", len(actualTopic.Queue), len(expectedTopic.Queue))
	}

	if (actualQueue[0].ID != expectedQueue[0].ID) || (actualQueue[0].Message != expectedQueue[0].Message) || (actualQueue[0].Topic != expectedQueue[0].Topic) {
		t.Errorf("Item struct payload is invalid.")
	}

}

func TestLength(t *testing.T) {
	var newTopic Topic
	actualTopic := newTopic.New("TestTopic")
	actualTopic.Insert("This is a test message")
	actualQueueLength := len(actualTopic.Queue)

	if actualQueueLength != 1 {
		t.Errorf("Expected queue length: %d got %d", len(actualTopic.Queue), 1)
	}

}

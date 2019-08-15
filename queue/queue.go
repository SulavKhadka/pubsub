package queue

import (
	"time"
)

// Topic the struct for storing message queue topics
type Topic struct {
	TopicName string `json:"topic_name"`
	Queue     []Item `json:"queue"`
}

// Item the payload struct for the queue
type Item struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
	Topic     string `json:"topic"`
}

//New creates and returns a new queue
func (q *Topic) New(topicName string) Topic {
	queue := make([]Item, 0, 10000)
	newTopic := Topic{topicName, queue}
	return newTopic
}

//Insert adds an item to the bottom of the queue
func (q *Topic) Insert(message string) {

	item := Item{}
	item.ID = len(q.Queue) + 1
	item.TimeStamp = time.Now().Format(time.RFC3339)
	item.Message = message
	item.Topic = q.TopicName

	q.Queue = append(q.Queue, item)
}

// Get retrieves the first item in the queue
func (q *Topic) Get() Item {

	if len(q.Queue) > 0 {
		item := q.Queue[0]
		q.Delete()
		return item
	}

	return Item{}
}

//Length gives back the length of the queue
func (q *Topic) Length() int {
	return len(q.Queue)
}

//Delete removes the top item from the queue
func (q *Topic) Delete() {
	q.Queue = q.Queue[1:len(q.Queue)]
}
